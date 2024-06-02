package audio

import (
	"os/exec"

	"codeberg.org/tomkoid/audstopper/internal/config"
)

type Player struct {
	Name          string
	PauseCommand  []string
	ResumeCommand []string
}

func listPlayers(config *config.Config) []Player {
	var players []Player

	// if these binaries exist on the running system
	if _, err := exec.LookPath("mpc"); err == nil && config.Mpc {
		players = append(players, Player{
			Name:          "mpc",
			PauseCommand:  []string{"mpc", "pause"},
			ResumeCommand: []string{"mpc", "play"},
		})
	}

	if _, err := exec.LookPath("playerctl"); err == nil && config.PlayerCtl {
		players = append(players, Player{
			Name:          "playerctl",
			PauseCommand:  []string{"playerctl", "pause"},
			ResumeCommand: []string{"playerctl", "resume"},
		})
	}

	return players
}
