package config

type Server struct {
	JWT     JWT     `json:"jwt" toml:"jwt" mapstructure:"jwt"`
	Zap     Zap     `json:"zap" toml:"zap" mapstructure:"zap"`
	Redis   Redis   `json:"redis" toml:"redis" mapstructure:"redis"`
	System  System  `json:"system" toml:"system" mapstructure:"system"`
	Mysql   Mysql   `json:"mysql" toml:"mysql" mapstructure:"mysql"`
	Mongodb Mongodb `json:"mongodb" toml:"mongodb" mapstructure:"mongodb"`
}
