package sql

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/buck119br/gears/pkg/gears/core/gtype/row"
	"github.com/buck119br/gears/pkg/gears/log"
)

func NewMySQLDatabase() Database {
	db := &mysqlDatabase{
		ready: false,
	}

	return db
}

type mysqlDatabase struct {
	mu    sync.Mutex
	ready bool
	db    *sql.DB

	c Config

	lastPingTimestamp int64
}

func (my *mysqlDatabase) Init(config Config) error {
	my.mu.Lock()
	defer my.mu.Unlock()

	if my.ready {
		return nil
	}

	my.c = config
	log.Infof("mysql database init with config: [%#v]", my.c)

	if err := my.instanceInit(); err != nil {
		return fmt.Errorf("instance init error: [%v]", err)
	}

	my.ready = true
	log.Infof("mysql database init finished")

	return nil
}

func (my *mysqlDatabase) Close() error {
	my.mu.Lock()
	defer my.mu.Unlock()

	if !my.ready {
		return nil
	}

	log.Infof("mysql database closing ...")
	return my.db.Close()
}

func (my *mysqlDatabase) Get() (*sql.DB, error) {
	my.mu.Lock()
	defer my.mu.Unlock()

	if !my.ready {
		return nil, ErrDatabaseNotReady
	}

	err := my.ping()
	if err == nil {
		return my.db, nil
	}
	log.Errorf("mysql database  ping error: [%v]", err)

	log.Warningf("mysql database instance re-init ...")
	my.ready = false
	if err = my.instanceInit(); err != nil {
		return nil, fmt.Errorf("instance re-init error: [%v]", err)
	}
	my.ready = true

	return my.db, nil
}

func (my *mysqlDatabase) instanceInit() error {
	instanceConfig := my.c.GetInstanceConfig()

	if instanceConfig.ConnectTimeout <= 1 {
		instanceConfig.ConnectTimeout = 2
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?readTimeout=%ds&writeTimeout=%ds&charset=%s",
		instanceConfig.Username, instanceConfig.Password,
		my.c.GetAddr(), my.c.GetPort(),
		my.c.GetInstance(),
		instanceConfig.ReadTimeout, instanceConfig.WriteTimeout,
		instanceConfig.Charset,
	)
	db, err := sql.Open(driverNameFromProtocol(my.c.GetProtocol()), dsn)
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

	my.db = db

	if err = my.ping(); err != nil {
		return fmt.Errorf("ping error: [%v]", err)
	}

	return nil
}

func (my *mysqlDatabase) ping() error {
	if !my.needPing() {
		return nil
	}

	log.Debugf("mysql database ping ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(my.c.GetInstanceConfig().ConnectTimeout)*time.Second)
	defer cancel()
	if err := my.db.PingContext(ctx); err != nil {
		return err
	}

	my.lastPingTimestamp = time.Now().Unix()

	return nil
}

func (my *mysqlDatabase) needPing() bool {
	if time.Now().Unix()-my.lastPingTimestamp > 10 {
		return true
	}
	return false
}

type MySQLConfig struct {
	Instance       string
	Addr           string
	AlterAddr      string
	Port           string
	InstanceConfig InstanceConfig
}

func (c MySQLConfig) GetModel() row.Model               { return row.Analytical }
func (c MySQLConfig) GetProtocol() row.Protocol         { return MySQL }
func (c MySQLConfig) GetInstance() string               { return c.Instance }
func (c MySQLConfig) GetAddr() string                   { return c.Addr }
func (c MySQLConfig) GetAlterAddr() string              { return c.AlterAddr }
func (c MySQLConfig) GetPort() string                   { return c.Port }
func (c MySQLConfig) GetInstanceConfig() InstanceConfig { return c.InstanceConfig }
