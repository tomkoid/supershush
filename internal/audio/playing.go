package audio

import (
	"strings"

	"codeberg.org/tomkoid/supershush/internal/tools"
)

func isPlaying(player string) (bool, []string) {
	playing := false
	players := []string{} 

	if player == "playerctl" {
		output, err := tools.RunCommand("playerctl", "status")
		if err != nil {
			return false, nil
		}

		if string(output) == "Playing\n" {
			playing = true
			players = append(players, player)
		}
	}
	if player == "mpc" {
		output, err := tools.RunCommand("mpc", "status")
		if err != nil {
			return false, nil
		}

		println(string(output))
		if strings.Contains(string(output), "[playing]") {
			playing = true
			players = append(players, player)
		}
	}

	// log.Printf("returned from isplaying: %v", players)
	return playing, players 
}
