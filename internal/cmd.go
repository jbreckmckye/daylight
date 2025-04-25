package internal

import (
	"fmt"
	"log"
	"time"
)

type exitCode int

const (
	exitOK  exitCode = 0
	exitERR exitCode = 1
)

// DaylightQuery is an obj representing a user's "request"
type DaylightQuery struct {
	Lat       float64
	Long      float64
	TZ        time.Location
	Date      time.Time
	IP        string
	Condensed bool
}

func Daylight() exitCode {
	defer func() {
		r := recover()
		if r != nil {
			fatal(fmt.Errorf("%s", r))
		}
	}()

	// Read configuration from user supplied arguments
	args := Arguments{}
	args.ReadFromCLI()
	config, err := args.Config()
	if err != nil {
		return fatal(err)
	}

	// If we are missing data, make request to IPInfo
	// This logic allows us to support offline use
	if config.MissingFields() != nil {
		ipInfo, err := FetchIPInfo()
		if err != nil {
			return fatal(err)
		}

		configFromAPI, err := ipInfo.Config()
		if err != nil {
			return fatal(err)
		}

		config.FillValues(configFromAPI)
	} else {
		fmt.Printf("\nOffline operation 💾\n")
	}

	// Sanity check - this should never fail, but...
	err = config.MissingFields()
	if err != nil {
		return fatal(err)
	}

	query := config.DaylightQuery()

  if query.Condensed {
		viewmodel := Condensed(query)
		formatted := viewmodel.FormatString()
		fmt.Print(formatted)

		return exitOK
	}

	viewmodel := Summary(query)
	rendered := render(viewmodel)
	fmt.Println(rendered)

	return exitOK
}

func fatal(err error) exitCode {
	log.Printf("%s\n", err)
	return exitERR
}
