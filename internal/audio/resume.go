package audio

import (
	"log"

	"codeberg.org/tomkoid/audstopper/internal/config"
	"codeberg.org/tomkoid/audstopper/internal/tools"
)

func resumeAudio(config *config.Config, givenPlayers []string) {
	log.Println("Resuming audio.")

	for _, player := range listPlayers(config) {
		for _, givenPlayer := range givenPlayers {
			if player.Name == givenPlayer {
				log.Printf("Resuming audio for %s\n", givenPlayer)
				_, err := tools.RunCommand(player.ResumeCommand[0], player.ResumeCommand[1:]...)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
