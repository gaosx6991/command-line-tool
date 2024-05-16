package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(MISSING_FILE_OPERAND)
		return
	}

	switch os.Args[1] {
	case "-h", "--help":
		fmt.Println(HELP_HINT)
	default:
		fmt.Println(UNRECOGNIZED_OPTION(os.Args[1]))
	}
}
