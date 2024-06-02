package audio

import (
	"log"

	"mrogalski.eu/go/pulseaudio"

	"codeberg.org/tomkoid/supershush/internal/config"
)

type mutedSink struct {
	Name string
	Players []string
}

func AudioMonitor(c *pulseaudio.Client, config *config.Config) {
	outputs, activeIndex, err := c.Outputs()
	if err != nil {
		c.Close()
		log.Fatal(err)
	}

	// Get the default sink (audio source)
	var defaultSink pulseaudio.Output = outputs[activeIndex]

	var initialSinkName string = defaultSink.CardID

	updateChan, err := c.Updates()
	if err != nil {
		c.Close()
		log.Fatal(err)
	}

	/// Muted sink
	var ms mutedSink

	log.Println("Starting audio monitoring.")

	// Monitor the default sink for changes
	for {
		<-updateChan

		// Get the new default sink
		outputs, activeIndex, err := c.Outputs()
		if err != nil {
			c.Close()
			log.Fatal(err)
		}

		defaultSink = outputs[activeIndex]

		// Check if the sink name has changed
		if defaultSink.CardID != initialSinkName {
			log.Printf(
				"Audio source changed from %s to %s\n",
				initialSinkName,
				defaultSink.CardID,
			)

			if config.Resume && ms.Name == defaultSink.CardID {
				log.Println(
					"Audio was muted, unmuted it.",
				)

				initialSinkName = defaultSink.CardID

				resumeAudio(config, ms.Players)
				ms.Name = ""
				ms.Players = []string{}
				continue
			}

			// stop audio if mpc or playerctl is running
			playingPlayers := stopAudio(config)
			log.Printf("Playing: %t\n", len(playingPlayers) != 0)
			if len(playingPlayers) != 0 {
				ms.Name = initialSinkName
				ms.Players = playingPlayers
			}
			// log.Printf("1: %s\n", ms.Name)
			// log.Printf("2: %s\n", defaultSink.CardID)

			initialSinkName = defaultSink.CardID
		}
	}
}
