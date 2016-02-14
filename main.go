package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("docker", "version")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
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
