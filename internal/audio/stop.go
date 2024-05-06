package audio

import (
	"log"
	"os/exec"

	"codeberg.org/tomkoid/audiochanger/internal/config"
	"codeberg.org/tomkoid/audiochanger/internal/tools"
)

func StopAudio(config *config.Config) {
	log.Println("Stopping audio")
	// if mpc binary exists on system
	if _, err := exec.LookPath("mpc"); err == nil && config.Mpc {
		err = tools.RunCommand("mpc", "pause")
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := exec.LookPath("playerctl"); err == nil && config.PlayerCtl {
		err = tools.RunCommand("playerctl", "pause")
		if err != nil {
			log.Println(err)
		}
	}
}
