package config

type Server struct {
	JWT    JWT    `json:"jwt" toml:"jwt"`
	Zap    Zap    `json:"zap" toml:"zap"`
	Redis  Redis  `json:"redis" toml:"redis"`
	System System `json:"system" toml:"system"`
	Mysql  Mysql  `json:"mysql" toml:"mysql"`
}
