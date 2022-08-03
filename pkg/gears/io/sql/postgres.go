package sql

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"

	"github.com/buck119br/gears/pkg/gears/core/gtype/row"
	"github.com/buck119br/gears/pkg/gears/log"
)

func NewPostgresDatabase() Database {
	db := &postgresDatabase{
		ready: false,
	}

	return db
}

type postgresDatabase struct {
	mu    sync.Mutex
	ready bool
	db    *sql.DB

	c Config

	lastPingTimestamp int64
}

func (p *postgresDatabase) Init(config Config) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.ready {
		return nil
	}

	p.c = config
	log.Infof("postgres database init with config: [%#v]", p.c)

	if err := p.instanceInit(); err != nil {
		return fmt.Errorf("instance init error: [%v]", err)
	}

	p.ready = true
	log.Infof("postgres database init finished")

	return nil
}

func (p *postgresDatabase) Get() (*sql.DB, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.ready {
		return nil, ErrDatabaseNotReady
	}

	err := p.ping()
	if err == nil {
		return p.db, nil
	}
	log.Errorf("postgres database ping error: [%v]", err)

	log.Warningf("postgres database instance re-init ...")
	p.ready = false
	if err := p.instanceInit(); err != nil {
		return nil, fmt.Errorf("instance re-init error: [%v]", err)
	}
	p.ready = true

	return p.db, nil
}

func (p *postgresDatabase) instanceInit() error {
	instanceConfig := p.c.GetInstanceConfig()

	if instanceConfig.ConnectTimeout <= 1 {
		instanceConfig.ConnectTimeout = 2
	}

	switch instanceConfig.SSLMode {
	case "disable", "allow", "prefer", "require", "verify-ca", "verify-full":
	default:
		return fmt.Errorf("invalid ssl mode:[%s]", instanceConfig.SSLMode)
	}

	db, err := sql.Open(driverNameFromProtocol(p.c.GetProtocol()),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?connect_timeout=%d&sslmode=%s",
			instanceConfig.Username, instanceConfig.Password,
			p.c.GetAddr(), p.c.GetPort(),
			p.c.GetInstance(),
			instanceConfig.ConnectTimeout, instanceConfig.SSLMode,
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

	p.db = db

	if err = p.ping(); err != nil {
		return fmt.Errorf("ping error: [%v]", err)
	}

	return nil
}

func (p *postgresDatabase) ping() error {
	if !p.needPing() {
		return nil
	}

	log.Debugf("postgres database ping ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.c.GetInstanceConfig().ConnectTimeout)*time.Second)
	defer cancel()
	if err := p.db.PingContext(ctx); err != nil {
		return err
	}

	p.lastPingTimestamp = time.Now().Unix()

	return nil
}

func (p *postgresDatabase) needPing() bool {
	if time.Now().Unix()-p.lastPingTimestamp > 10 {
		return true
	}
	return false
}

type PostgresConfig struct {
	Instance       string
	Addr           string
	Port           string
	InstanceConfig InstanceConfig
}

func (c PostgresConfig) GetModel() row.Model               { return row.Relational }
func (c PostgresConfig) GetProtocol() row.Protocol         { return PostgreSQL }
func (c PostgresConfig) GetInstance() string               { return c.Instance }
func (c PostgresConfig) GetAddr() string                   { return c.Addr }
func (c PostgresConfig) GetAlterAddr() string              { return "" }
func (c PostgresConfig) GetPort() string                   { return c.Port }
func (c PostgresConfig) GetInstanceConfig() InstanceConfig { return c.InstanceConfig }
