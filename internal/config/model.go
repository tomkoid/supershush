package config

type Config struct {
	Resume    bool `toml:"resume"`
	Mpc       bool `toml:"mpc"`
	PlayerCtl bool `toml:"playerctl"`
}
