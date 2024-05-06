package config

type Config struct {
	Mpc       bool `toml:"mpc"`
	PlayerCtl bool `toml:"playerctl"`
}
