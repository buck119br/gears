package sql

import (
	"errors"
	"fmt"

	"github.com/buck119br/gears/pkg/gears/core/gtype/row"
)

const (
	DefaultMaxNumConnOpen = 5
	DefaultMaxNumConnIdle = 2
)

var (
	ClickHouseSQL = row.NewProtocol("ClickHouseSQL")
	MySQL         = row.NewProtocol("MySQL")
	PostgreSQL    = row.NewProtocol("PostgreSQL")
)

var (
	ErrDatabaseNotReady = errors.New("not ready")
)

func driverNameFromProtocol(protocol row.Protocol) string {
	switch protocol {
	case ClickHouseSQL:
		return "clickhouse"
	case PostgreSQL:
		return "postgres"
	default:
		panic(fmt.Errorf("invalid protocol: [%s]", protocol))
	}
}
