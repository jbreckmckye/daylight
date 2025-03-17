package cmd2

type exitCode int

const (
	exitOK  exitCode = 0
	exitERR exitCode = 1
)

func Daylight() exitCode {
  return 0
}


func parseFlags() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")

	flag.Parse()
	fmt.Println("word:", *wordPtr)
  fmt.Println("numb:", *numbPtr)
}
