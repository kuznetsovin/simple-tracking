package main

import (
	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
	"simple-tracking/backend/utils"
)

type Config struct {
	ServiceName string          `toml:"service_name"`
	LogLevel    string          `toml:"log_level"`
	Addr        string          `toml:"addr"`
	DbConn      string          `toml:"db_conn"`
	Broker      utils.BrokerCfg `toml:"broker"`
}

func (c *Config) GetLogLever() log.Lvl {
	var lvl log.Lvl

	switch c.LogLevel {
	case "DEBUG":
		lvl = log.DEBUG
		break
	case "INFO":
		lvl = log.INFO
		break
	case "WARN":
		lvl = log.WARN
		break
	case "ERROR":
		lvl = log.ERROR
		break
	default:
		lvl = log.INFO
	}
	return lvl
}

func (c *Config) Load(path string) error {
	_, err := toml.DecodeFile(path, c)
	return err
}
