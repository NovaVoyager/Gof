package config

type Mongodb struct {
	Host string `json:"host" toml:"host" mapstructure:"host"`
	Port int    `json:"port" toml:"port" mapstructure:"port"`
}
