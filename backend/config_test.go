package main

import (
	"github.com/stretchr/testify/assert"
	"simple-tracking/backend/utils"
	"testing"
)

func TestConfig_Load(t *testing.T) {
	c := Config{}
	if assert.NoError(t, c.Load("config.toml")) {
		assert.Equal(t, Config{
			ServiceName: "map_ws",
			LogLevel:    "DEBUG",
			Addr:        "localhost:8081",
			DbConn:      "postgresql://localhost:5432/skolkovo?sslmode=disable",
			Broker: utils.BrokerCfg{
				"amqp://guest:guest@localhost:5672/",
				"new_ws",
				1000,
				0,
				true,
			},
		}, c)
	}
}
