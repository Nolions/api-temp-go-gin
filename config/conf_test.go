package config

import (
	"github.com/redpkg/formula/db"
	"github.com/redpkg/formula/log"
	"github.com/redpkg/formula/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
		// Add test cases.
		{
			name: "app",
			args: "app.conf.example.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotConf, err := Init(tt.args)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("gotConf = %+v", gotConf)
			if tt.name == "app" {
				assertConfEqual(t, newAppConf(), gotConf)
			}
		})
	}
}

func assertConfEqual(t *testing.T, expected, actual *Conf) {
	assert.Equal(t, expected, actual)
}

func newAppConf() *Conf {
	return &Conf{
		Project: "api-temp",
		App: App{
			Addr:         "8999",
			Mode:         "debug",
			ReadTimeout:  "10s",
			WriteTimeout: "10s",
		},
		Log: log.Config{
			Level:   "debug",
			Console: false,
		},
		DB: db.Config{
			Driver:   "mysql",
			Database: "platform",
			Master: db.ConfigNode{
				Username: "root",
				Password: "root",
				Address:  "localhost:3306",
			},
			Slave: db.ConfigNode{
				Username: "root",
				Password: "root",
				Address:  "localhost:3306",
			},
			DialTimeout:     "10s",
			ReadTimeout:     "30s",
			WriteTimeout:    "60s",
			DBTimezone:      "UTC",
			AppTimezone:     "UTC",
			ConnMaxLifeTime: "0s",
			MaxIdleConns:    2,
			MaxOpenConns:    0,
		},
		Redis: redis.Config{
			Address:  "localhost:6379",
			Password: "",
			DB:       0,
		},
	}
}
