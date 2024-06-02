package audio

import (
	"log"

	"codeberg.org/tomkoid/audstopper/internal/config"
	"codeberg.org/tomkoid/audstopper/internal/tools"
)

/// returns an array of players that were playing before stopped
func stopAudio(config *config.Config) []string {
	log.Println("Stopping audio.")

	var players []string
	for _, player := range listPlayers(config) {
		currentPlaying, currentPlayers := isPlaying(player.Name)

		if currentPlaying {
			for _, playingPlayer := range currentPlayers {
				players = append(players, playingPlayer)
			}
		}

		_, err := tools.RunCommand(player.PauseCommand[0], player.PauseCommand[1:]...)
		if err != nil {
			log.Println(err)
		}
	}

	return players 
}
