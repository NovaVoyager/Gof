package config

type System struct {
	Host string `json:"host" toml:"host" mapstructure:"host"`
	Mode string `json:"mode" toml:"mode" mapstructure:"mode"`
}
