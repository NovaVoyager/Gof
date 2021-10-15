package config

type Mysql struct {
	Host         string `json:"host" toml:"host"`
	Port         int    `json:"port" toml:"port"`
	Config       string `json:"config" toml:"config"`
	DbName       string `json:"db_name" toml:"db_name"`
	Username     string `json:"username" toml:"username"`
	Password     string `json:"password" toml:"password"`
	MaxIdleConns int    `json:"max_idle_conns" toml:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" toml:"max_open_conns"`
	LogMode      bool   `json:"log_mode" toml:"log_mode"`
	LogZap       string `json:"log_zap" toml:"log_zap"`
}
