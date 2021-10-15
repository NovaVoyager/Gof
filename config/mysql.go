package config

type Mysql struct {
	Host         string `json:"host" toml:"host" mapstructure:"host"`
	Port         int    `json:"port" toml:"port" mapstructure:"port"`
	Config       string `json:"config" toml:"config" mapstructure:"config"`
	DbName       string `json:"db_name" toml:"db_name" mapstructure:"db_name"`
	Username     string `json:"username" toml:"username" mapstructure:"username"`
	Password     string `json:"password" toml:"password" mapstructure:"password"`
	MaxIdleConns int    `json:"max_idle_conns" toml:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" toml:"max_open_conns" mapstructure:"max_open_conns"`
	LogMode      bool   `json:"log_mode" toml:"log_mode" mapstructure:"log_mode"`
	LogZap       string `json:"log_zap" toml:"log_zap" mapstructure:"log_zap"`
}
