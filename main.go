package main

import (
	"log"

	"mrogalski.eu/go/pulseaudio"

	"codeberg.org/tomkoid/supershush/internal/audio"
	"codeberg.org/tomkoid/supershush/internal/config"
	"codeberg.org/tomkoid/supershush/internal/tools"
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
