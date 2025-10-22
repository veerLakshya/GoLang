package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func subcommands() {

	stringFlag := flag.String("user", "guest", "Name of the user")
	flag.StringVar(stringFlag, "user", "John", "Name of user")
	sub1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	sub2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := sub1.Bool("processing", false, "command processing status")
	secondFlag := sub1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := sub2.String("language", "Go", "Enter Your Language")

	if len(os.Args) < 2 {
		fmt.Println("This Program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		sub1.Parse(os.Args[2:])
		fmt.Println("sub1:")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		sub2.Parse(os.Args[2:])
		fmt.Println("subCommand2: ")
		fmt.Println("language: ", flagsc2)
	default:
		fmt.Println("no subCommands entered")
		os.Exit(1)
	}
}
