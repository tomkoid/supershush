package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"mrogalski.eu/go/pulseaudio"

	"codeberg.org/tomkoid/audiochanger/internal/config"
)

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func stopAudio(config *config.Config) {
	log.Println("Stopping audio")
	// if mpc binary exists on system
	if _, err := exec.LookPath("mpc"); err == nil && config.Mpc {
		err = runCommand("mpc", "pause")
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := exec.LookPath("playerctl"); err == nil && config.PlayerCtl {
		err = runCommand("playerctl", "pause")
		if err != nil {
			log.Println(err)
		}
	}
}

func handleCleanup(pulseClient *pulseaudio.Client) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	println("Shutting down...")

	// Run Cleanup
	pulseClient.Close()

	os.Exit(1)
}

func main() {
	config := config.GetConfig()

	// test()

	// Create a PulseAudio context
	c, err := pulseaudio.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	go handleCleanup(c)

	defer c.Close()

	outputs, activeIndex, err := c.Outputs()
	if err != nil {
		log.Fatal(err)
	}

	// Get the default sink (audio source)
	var defaultSink pulseaudio.Output = outputs[activeIndex]

	// Store the initial sink name
	var initialSinkName string = defaultSink.CardID

	updateChan, err := c.Updates()
	if err != nil {
		log.Fatal(err)
	}

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
			fmt.Printf(
				"Audio source changed from %s to %s\n",
				initialSinkName,
				defaultSink.CardID,
			)

			// stop audio if mpc or playerctl is running
			stopAudio(&config)

			initialSinkName = defaultSink.CardID
		}
	}
}
