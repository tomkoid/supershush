package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/jfreymuth/pulse"

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

func handleCleanup(pulseClient *pulse.Client) {
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

	// Create a PulseAudio context
	c, err := pulse.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	go handleCleanup(c)

	defer c.Close()

	// Get the default sink (audio source)
	defaultSink, err := c.DefaultSink()
	if err != nil {
		log.Fatal(err)
	}

	// Store the initial sink name
	initialSinkName := defaultSink.Name()

	// Monitor the default sink for changes
	for {
		// Get the new default sink
		defaultSink, err = c.DefaultSink()
		if err != nil {
			log.Fatal(err)
		}

		// Check if the sink name has changed
		if defaultSink.Name() != initialSinkName {
			fmt.Printf(
				"Audio source changed from %s to %s\n",
				initialSinkName,
				defaultSink.Name(),
			)

			// stop audio if mpc or playerctl is running
			stopAudio(&config)

			initialSinkName = defaultSink.Name()
		}

		time.Sleep(time.Duration(config.PollRateMs) * time.Millisecond)
	}
}
