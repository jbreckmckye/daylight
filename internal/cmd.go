package internal

import (
	"fmt"
	"log"
	"time"
)

type ExitCode int

const (
	exitOK  ExitCode = 0
	exitERR ExitCode = 1
)

// DaylightQuery is an obj representing a user's "request"
type DaylightQuery struct {
	Lat       float64
	Long      float64
	TZ        time.Location
	Date      time.Time
	IP        string
	Condensed bool
	Json      bool
}

func Daylight(version string) ExitCode {
	defer func() {
		r := recover()
		if r != nil {
			fatal(fmt.Errorf("%s", r))
		}
	}()

	// Read configuration from user supplied arguments
	args := Arguments{}
	args.ReadFromCLI()

	if args.Version != nil && *args.Version {
		fmt.Println(version)
		return exitOK
	}
	
	config, err := args.Config()
	if err != nil {
		return fatal(err)
	}

	// If we are missing data, make request to IPInfo
	// Otherwise we can run in "offline mode"
	offline := config.MissingFields() == nil
	if !offline {
		ipInfo, err := FetchIPInfo()
		if err != nil {
			return fatal(err)
		}

		configFromAPI, err := ipInfo.Config()
		if err != nil {
			return fatal(err)
		}

		config.FillValues(configFromAPI)
	}

	// Sanity check - this should never fail, but...
	err = config.MissingFields()
	if err != nil {
		return fatal(err)
	}

	query := config.DaylightQuery()

	if query.Json {
		viewmodel := JsonSummary(query)
		formatted, err := viewmodel.FormatString()
		if err != nil {
			return fatal(err)
		}
		fmt.Println(formatted)

		return exitOK
	}

	if query.Condensed {
		viewmodel := Condensed(query)
		formatted := viewmodel.FormatString()
		fmt.Print(formatted)

		return exitOK
	}

	if offline {
		fmt.Print(Offline())
	}
	viewmodel := Summary(query)
	formatted := viewmodel.FormatString()
	fmt.Print(formatted)

	return exitOK
}

func fatal(err error) ExitCode {
	log.Printf("%s\n", err)
	return exitERR
}
