package config

type JWT struct {
	SigningKey string `json:"signing_key" toml:"signing_key" mapstructure:"signing_key"`
	ExpireTime string `json:"expire_time" toml:"expire_time" mapstructure:"expire_time"`
	BufferTime string `json:"buffer_time" toml:"buffer_time" mapstructure:"buffer_time"`
}
