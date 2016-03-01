package main

import (
	"github.com/katherinealbany/rodentia/logger"
	"os"
	"os/exec"
)

var log = logger.New("main")

func main() {
	cmd := exec.Command("docker", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Debug("Starting...")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	log.Info("Running...")

	log.Info("Waiting...")
	err := cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
