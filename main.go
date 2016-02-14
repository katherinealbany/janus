package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("docker", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	log.Println("Running...")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Waiting...")
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Finished with error: %v", err)
}
