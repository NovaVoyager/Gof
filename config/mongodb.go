package config

type Mongodb struct {
	Host     string `json:"host" toml:"host" mapstructure:"host"`
	Port     int    `json:"port" toml:"port" mapstructure:"port"`
	Username string `json:"username" toml:"username" mapstructure:"username"`
	Password string `json:"password" toml:"password" mapstructure:"password"`
}
