package utils

type BrokerCfg struct {
	Conn          string `toml:"conn"`
	Queue         string `toml:"queue"`
	PrefetchCount int    `toml:"prefetch_count"`
	PrefetchSize  int    `toml:"prefetch_size"`
	QosGlobal     bool   `toml:"qos_global"`
}
