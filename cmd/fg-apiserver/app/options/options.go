package options

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// MySQLOptions defines options for mysql database
type MySQLOptions struct {
	Addr                  string        `json:"addr,omitempty"`
	Username              string        `json:"username,omitempty"`
	Password              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
}

// NewMySQLOptions create a zero value instance
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:                  "127.0.0.1:3306",
		Username:              "onex",
		Password:              "onex(#)666",
		Database:              "onex",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

type ServerOptions struct {
	MySQLOptions *MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: NewMySQLOptions(),
	}
}

// Validate checks ServerOptions
func (o *ServerOptions) Validate() error {

	if o.MySQLOptions.Addr == "" {
		return fmt.Errorf("MySQL server address cannot be empty")
	}
	// check the address is like host:port
	host, portStr, err := net.SplitHostPort(o.MySQLOptions.Addr)
	if err != nil {
		return fmt.Errorf("Invalid MySQL address format '%s': %w", o.MySQLOptions.Addr, err)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("Invalid MySQL port: %s", portStr)
	}

	if host == "" {
		return fmt.Errorf("MySQL hostname cannot be empty")
	}

	if o.MySQLOptions.Username == "" {
		return fmt.Errorf("MySQL username cannot be empty")
	}

	if o.MySQLOptions.Password == "" {
		return fmt.Errorf("MySQL password cannot be empty")
	}

	if o.MySQLOptions.Database == "" {
		return fmt.Errorf("MySQL database name cannot be empty")
	}

	// Check connection pool
	if o.MySQLOptions.MaxIdleConnections <= 0 {
		return fmt.Errorf("MySQL max idle connections must be greater than 0")
	}

	if o.MySQLOptions.MaxOpenConnections <= 0 {
		return fmt.Errorf("MySQL max open connections must be greater than 0")
	}

	if o.MySQLOptions.MaxIdleConnections > o.MySQLOptions.MaxOpenConnections {
		return fmt.Errorf("MySQL max idle connections cannot be greater than max open connections")
	}

	if o.MySQLOptions.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("MySQL max connection lifetime must be greater than 0")
	}

	return nil
}
