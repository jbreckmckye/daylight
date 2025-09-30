package main

import (
	"log"
	"os"

	_ "time/tzdata"

	"github.com/jbreckmckye/daylight/internal"
)

// Set by GoReleaser via ldflags
var version = "development"

func main() {
	log.SetPrefix("[daylight] ")
	log.SetFlags(0)

	code := internal.Daylight(version)

	os.Exit(int(code))
}
