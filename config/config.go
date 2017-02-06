package config

import (
	"fmt"
	"net/url"

	"github.com/Sirupsen/logrus"
)

type Config struct {
	BindPort     int    `mapstructure:"PORT" yaml:"-"`
	BindHost     string `mapstructure:"HOST" yaml:"-"`
	SystemSecret string `mapstructure:"SYSTEM_SECRET" yaml:"-"`
	DatabaseURL  string `mapstructure:"DATABASE_URL" yaml:"-"`
	LogLevel     string `yaml:"-"`

	context *Context `yaml:"-"`
}

func Load() {
	fmt.Println("loading config")
}

func (c *Config) Context() *Context {
	if c.context != nil {
		return c.context
	}

	u, err := url.Parse(c.DatabaseURL)
	if err != nil {
		logrus.Fatalf("Could not parse DATABASE_URL: %s", err)
	}

	var connection = interface{}(nil)
	switch u.Scheme {
	case "mysql":
		connection = &SQLConnection{URL: u}
		break
	default:
		logrus.Fatalf("Unkown DSN %s in DATABASE_URL: %s", u.Scheme, c.DatabaseURL)
	}

	c.context = &Context{
		Connection: connection,
	}

	return c.context
}

func (c *Config) getAddress() string {
	return fmt.Sprintf("%s:%d", c.BindHost, c.BindPort)
}
