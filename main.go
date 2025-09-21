package main

import (
	"log"
	"os"

	_ "time/tzdata"

	"github.com/jbreckmckye/daylight/internal"
)

func main() {
	log.SetPrefix("[daylight] ")
	log.SetFlags(0)

	code := internal.Daylight()

	os.Exit(int(code))
}
