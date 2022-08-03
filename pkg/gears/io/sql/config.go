package sql

import (
	"github.com/buck119br/gears/pkg/gears/core/gtype/row"
)

type Config interface {
	GetModel() row.Model
	GetProtocol() row.Protocol
	GetInstance() string
	GetAddr() string
	GetAlterAddr() string
	GetPort() string
	GetInstanceConfig() InstanceConfig
}

type InstanceConfig struct {
	Username       string
	Password       string
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
	MaxOpen        int
	MaxIdle        int

	// PostgreSQL exclusive
	SSLMode string
}
