package main

import (
	"log"
	"os"

	"github.com/jbreckmckye/daylight/internal/new"
)

func main() {
	log.SetPrefix("[daylight] ")
	log.SetFlags(0)

	code := new.Daylight()

	os.Exit(int(code))
}
