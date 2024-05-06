package main

import (
	"log"

	"mrogalski.eu/go/pulseaudio"

	"codeberg.org/tomkoid/audiochanger/internal/audio"
	"codeberg.org/tomkoid/audiochanger/internal/config"
	"codeberg.org/tomkoid/audiochanger/internal/tools"
)

func main() {
	config := config.GetConfig()

	// Create a PulseAudio context
	c, err := pulseaudio.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Handle program termination
	go tools.HandleCleanup(c)

	defer c.Close()

	outputs, activeIndex, err := c.Outputs()
	if err != nil {
		log.Fatal(err)
	}

	// Get the default sink (audio source)
	var defaultSink pulseaudio.Output = outputs[activeIndex]

	var initialSinkName string = defaultSink.CardID

	updateChan, err := c.Updates()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting audio monitoring.")

	// Monitor the default sink for changes
	for {
		<-updateChan

		// Get the new default sink
		outputs, activeIndex, err := c.Outputs()
		if err != nil {
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
			audio.StopAudio(&config)

			initialSinkName = defaultSink.CardID
		}
	}
}
