package main

import (
	"log"

	"mrogalski.eu/go/pulseaudio"

	"codeberg.org/tomkoid/audstopper/internal/audio"
	"codeberg.org/tomkoid/audstopper/internal/config"
	"codeberg.org/tomkoid/audstopper/internal/tools"
)

func main() {
	cfg := config.GetConfig()

	// Create a PulseAudio context
	c, err := pulseaudio.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Handle program termination
	go tools.HandleCleanup(c)

	// Start audio monitoring
	audio.AudioMonitor(c, &cfg)

	defer c.Close()
}
