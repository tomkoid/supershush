package audio

import (
	"log"

	"mrogalski.eu/go/pulseaudio"

	"codeberg.org/tomkoid/audstopper/internal/config"
)

type mutedSink struct {
	Name string
	Playing
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
	log.Println(ms.Name)

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

			if ms.Name == defaultSink.CardID {
				log.Println(
					"Audio was muted, so unmuted it.!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!",
				)

				initialSinkName = defaultSink.CardID

				ms.Name = ""
				resumeAudio(config, ms.Players)
				continue
			}

			// stop audio if mpc or playerctl is running
			playing := stopAudio(config)
			log.Printf("Playing: %t\n", playing.Playing)
			if playing.Playing {
				ms.Name = initialSinkName
				ms.Players = playing.Players
			}
			log.Printf("1: %s\n", ms.Name)
			log.Printf("2: %s\n", defaultSink.CardID)

			initialSinkName = defaultSink.CardID
		}
	}
}
