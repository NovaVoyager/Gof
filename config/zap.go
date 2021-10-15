package config

type Zap struct {
	Level         string `json:"level" toml:"level" mapstructure:"level"`
	Format        string `json:"format" toml:"format" mapstructure:"format"`
	Prefix        string `json:"prefix" toml:"prefix" mapstructure:"prefix"`
	Director      string `json:"director" toml:"director" mapstructure:"director"`
	LinkName      string `json:"link_name" toml:"link_name" mapstructure:"link_name"`
	ShowLine      bool   `json:"show_line" toml:"show_line" mapstructure:"show_line"`
	EncodeLevel   string `json:"encode_level" toml:"encode_level" mapstructure:"encode_level"`
	StacktraceKey string `json:"stacktrace_key" toml:"stacktrace_key" mapstructure:"stacktrace_key"`
	LogInConsole  bool   `json:"log_in_console" toml:"log_in_console" mapstructure:"log_in_console"`
}
