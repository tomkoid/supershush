package audio

import (
	"strings"

	"codeberg.org/tomkoid/audstopper/internal/tools"
)

func isPlaying(player string) (bool, *string) {
	if player == "playerctl" {
		output, err := tools.RunCommand("playerctl", "status")
		if err != nil {
			return false, nil
		}

		if string(output) == "Playing\n" {
			println("playerctl playing!!!!!!!!!!!!!!!!!!!!!!!!!")
			return true, &player
		}
	}
	if player == "mpc" {
		output, err := tools.RunCommand("mpc", "status")
		if err != nil {
			return false, nil
		}

		println(string(output))
		if strings.Contains(string(output), "[playing]") {
			println("mpc playing!!!!!!!!!!!!!!!!!!!!!!!!!")
			return true, &player
		}
	}

	return false, nil
}
