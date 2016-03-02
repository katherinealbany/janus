package main

import (
	"flag"
	"github.com/katherinealbany/rodentia/logger"
	"os"
	"os/exec"
)

var log = logger.New("main")
var dir string

func main() {
	flag.StringVar(&dir, "dir", ".", "supply it")
	flag.Parse()
	log.Info(dir)

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
