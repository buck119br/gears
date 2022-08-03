package sql

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"

	"github.com/buck119br/gears/pkg/gears/core/gtype/row"
	"github.com/buck119br/gears/pkg/gears/log"
)

func NewClickHouseDatabase() Database {
	db := &clickHouseDatabase{
		ready: false,
	}

	return db
}

type clickHouseDatabase struct {
	mu    sync.Mutex
	ready bool
	db    *sql.DB

	c Config

	lastPingTimestamp int64
}

func (ch *clickHouseDatabase) Init(config Config) error {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if ch.ready {
		return nil
	}

	ch.c = config
	log.Infof("clickhouse database init with config: [%#v]", ch.c)

	if err := ch.instanceInit(); err != nil {
		return fmt.Errorf("instance init error: [%v]", err)
	}

	ch.ready = true
	log.Infof("clickhouse database init finished")

	return nil
}

func (ch *clickHouseDatabase) Get() (*sql.DB, error) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if !ch.ready {
		return nil, ErrDatabaseNotReady
	}

	err := ch.ping()
	if err == nil {
		return ch.db, nil
	}
	log.Errorf("clickhouse database  ping error: [%v]", err)

	log.Warningf("clickhouse database instance re-init ...")
	ch.ready = false
	if err = ch.instanceInit(); err != nil {
		return nil, fmt.Errorf("instance re-init error: [%v]", err)
	}
	ch.ready = true

	return ch.db, nil
}

func (ch *clickHouseDatabase) instanceInit() error {
	instanceConfig := ch.c.GetInstanceConfig()

	if instanceConfig.ConnectTimeout <= 1 {
		instanceConfig.ConnectTimeout = 2
	}
	db, err := sql.Open(driverNameFromProtocol(ch.c.GetProtocol()),
		fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s&read_timeout=%d&write_timeout=%d&alt_hosts=%s",
			ch.c.GetAddr(), ch.c.GetPort(),
			instanceConfig.Username, instanceConfig.Password,
			ch.c.GetInstance(),
			instanceConfig.ReadTimeout, instanceConfig.WriteTimeout,
			ch.c.GetAlterAddr(),
		),
	)
	if err != nil {
		return fmt.Errorf("open error: [%v]", err)
	}

	if instanceConfig.MaxOpen <= 0 {
		instanceConfig.MaxOpen = DefaultMaxNumConnOpen
	}
	if instanceConfig.MaxIdle <= 0 {
		instanceConfig.MaxIdle = DefaultMaxNumConnIdle
	}
	db.SetMaxOpenConns(instanceConfig.MaxOpen)
	db.SetMaxIdleConns(instanceConfig.MaxIdle)

	ch.db = db

	if err = ch.ping(); err != nil {
		return fmt.Errorf("ping error: [%v]", err)
	}

	return nil
}

func (ch *clickHouseDatabase) ping() error {
	if !ch.needPing() {
		return nil
	}

	log.Debugf("clickhouse database ping ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ch.c.GetInstanceConfig().ConnectTimeout)*time.Second)
	defer cancel()
	if err := ch.db.PingContext(ctx); err != nil {
		return err
	}

	ch.lastPingTimestamp = time.Now().Unix()

	return nil
}

func (ch *clickHouseDatabase) needPing() bool {
	if time.Now().Unix()-ch.lastPingTimestamp > 10 {
		return true
	}
	return false
}

type ClickHouseConfig struct {
	Instance       string
	Addr           string
	AlterAddr      string
	Port           string
	InstanceConfig InstanceConfig
}

func (c ClickHouseConfig) GetModel() row.Model               { return row.Analytical }
func (c ClickHouseConfig) GetProtocol() row.Protocol         { return ClickHouseSQL }
func (c ClickHouseConfig) GetInstance() string               { return c.Instance }
func (c ClickHouseConfig) GetAddr() string                   { return c.Addr }
func (c ClickHouseConfig) GetAlterAddr() string              { return c.AlterAddr }
func (c ClickHouseConfig) GetPort() string                   { return c.Port }
func (c ClickHouseConfig) GetInstanceConfig() InstanceConfig { return c.InstanceConfig }
