package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
