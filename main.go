package main

import (
	"github.com/katherinealbany/rodentia/logger"
	"os"
	"os/exec"
)

var log = logger.New("main")

func main() {
	args := os.Args[1:]
	log.Info(args)

	cmd := exec.Command("docker", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Debug("Starting...")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	log.Info("Running...")
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
