package audio

import (
	"log"

	"codeberg.org/tomkoid/audstopper/internal/config"
	"codeberg.org/tomkoid/audstopper/internal/tools"
)

type Playing struct {
	Playing bool
	Players []string
}

// returns if audio was playing
func stopAudio(config *config.Config) Playing {
	log.Println("Stopping audio.")

	var playing Playing
	for _, player := range listPlayers(config) {
		currentPlaying, currentPlayer := isPlaying(player.Name)

		if currentPlaying {
			playing.Playing = true

			playing.Players = append(playing.Players, *currentPlayer)
		}

		_, err := tools.RunCommand(player.PauseCommand[0], player.PauseCommand[1:]...)
		if err != nil {
			log.Println(err)
		}
	}

	return playing
}
