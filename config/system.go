package config

type System struct {
	Host string `json:"host" toml:"host"`
	Mode string `json:"mode" toml:"mode"`
}
