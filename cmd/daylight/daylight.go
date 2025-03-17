package main

import (
	"log"
	"os"

	"github.com/jbreckmckye/daylight/internal/cmd"
)

func main() {
	log.SetPrefix("[daylight]")
	code := cmd.Daylight()
	os.Exit(int(code))
}
