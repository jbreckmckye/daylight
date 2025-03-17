package main

import (
	"os"

	"github.com/jbreckmckye/daylight/internal/cmd"
)

func main() {
	code := cmd.Daylight()
	os.Exit(code)
}
