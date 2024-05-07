package audio

import (
	"log"

	"mrogalski.eu/go/pulseaudio"

	"codeberg.org/tomkoid/audstopper/internal/config"
)

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

			// stop audio if mpc or playerctl is running
			stopAudio(config)

			initialSinkName = defaultSink.CardID
		}
	}
}
