package config

import (
	"net/url"

	"strings"
	"time"

	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/arisro/talos/utils"
	"github.com/jmoiron/sqlx"
)

type SQLConnection struct {
	db  *sqlx.DB
	URL *url.URL
}

func (c *SQLConnection) GetDatabase() *sqlx.DB {
	if c.db != nil {
		return c.db
	}

	var err error
	if err = utils.Retry(time.Second*15, time.Minute*2, func() error {
		logrus.Info("Connection with %s", c.URL.Scheme+"://*:*@"+c.URL.Host+c.URL.Path+"?"+c.URL.RawQuery)
		u := c.URL.String()
		if c.URL.Scheme == "mysql" {
			u = strings.Replace(u, "mysql://", "", -1)
		}

		if c.db, err = sqlx.Open(c.URL.Scheme, u); err != nil {
			return fmt.Errorf("Could not connect to SQL: %s", err)
		} else if err := c.db.Ping(); err != nil {
			return fmt.Errorf("Could not connect to SQL: %s", err)
		}

		logrus.Infof("Connected to SQL!")
		return nil
	}); err != nil {
		logrus.Fatalf("Could not connect to SQL: %s", err)
	}

	return c.db
}
