package rdb

type ConnectionConfig struct {
	Driver   string `toml:"driver"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	UserName string `toml:"user_name"`
	Password string `toml:"password"`
	Schema   string `toml:"schema"`
}
