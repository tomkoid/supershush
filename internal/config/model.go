package config

type Config struct {
	PollRateMs int  `toml:"poll_rate_ms"`
	Mpc        bool `toml:"mpc"`
	PlayerCtl  bool `toml:"playerctl"`
}
