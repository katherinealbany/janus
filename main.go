package main

import (
	"flag"
	"github.com/katherinealbany/rodentia/logger"
	"os"
	"os/exec"
)

var (
	log = logger.New("main")

	dir     string
	repo    string
	build   string
	stable  string
	release string
	force   string
	push    string
)

func init() {
	flag.StringVar(&dir, "dir", ".", "build directory")
	flag.StringVar(&repo, "repo", "katherinealbany", "docker registry repository")
	flag.StringVar(&build, "build", "latest", "build tag")
	flag.StringVar(&stable, "stable", "stable", "stable build tag")
	flag.StringVar(&release, "release", "v1.0.0", "release version")
	flag.StringVar(&force, "force", "false", "force the matter!")
	flag.StringVar(&push, "push", "true", "push after build")
}

func main() {
	log.Info("Parsing...")
	flag.Parse()

	log.Debug("dir     =", dir)
	log.Debug("repo    =", repo)
	log.Debug("build   =", build)
	log.Debug("stable  =", stable)
	log.Debug("release =", release)
	log.Debug("force   =", force)
	log.Debug("push    =", push)

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
