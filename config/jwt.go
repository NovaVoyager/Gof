package config

type JWT struct {
	SigningKey string `json:"signing_key" toml:"signing_key"`
	ExpireTime string `json:"expire_time" toml:"expire_time"`
	BufferTime string `json:"buffer_time" toml:"buffer_time"`
}
