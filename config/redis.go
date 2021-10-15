package config

type Redis struct {
	Host     string `json:"host" toml:"host" mapstructure:"host"`
	Port     int    `json:"port" toml:"port" mapstructure:"port"`
	Password string `json:"password" toml:"password" mapstructure:"password"`
}
