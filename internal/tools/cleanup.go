package tools

import (
	"os"
	"os/signal"
	"syscall"

	"mrogalski.eu/go/pulseaudio"
)

func HandleCleanup(pulseClient *pulseaudio.Client) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	println("Shutting down...")

	// Run Cleanup
	pulseClient.Close()

	os.Exit(1)
}
