package main

import (
	//"flag"
	//"os"
	//"github.com/jessevdk/go-flags"
	"fmt"
	"github.com/alexflint/go-arg"
)

// func main() {
// 	wordPtr := flag.String("word", "foo", "a string")
// 	numbPtr := flag.Int("numb", 42, "an int")

// 	flag.Parse()
// 	fmt.Println("word:", *wordPtr)
//   fmt.Println("numb:", *numbPtr)
// }

// var opts struct {
// 	Name string `short:"n" long:"name" description:"Your name" required:"true"`
// 	Latitude float64 `long:"lat" description:"Latitude"`
// 	Longitude float64 `long:"long" description:"Longitude"`
// }

// func main() {
// 	fmt.Printf("ARGS = %v\n", os.Args)

// 	parser := flags.NewParser(&opts, flags.Default)
//   // to be honest, I'm finding that the errors make more sense with goarg
//   unparsed, err := parser.Parse()
//   if err != nil {
// 		fmt.Printf("some error = %v\n", err)
// 	}

// 	fmt.Printf("unparsed = %v\n", unparsed)
// 	fmt.Printf("opts = %v\n", opts)
// 	fmt.Printf("latitude = %v ; was set %t \n", opts.Latitude, parser.FindOptionByLongName("lat").IsSet())

// 	return
// }

func mainr() {
	var args struct {
		Foo string
		Bar *float64 // pointer lets us check if the value was set
	}
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar)
}