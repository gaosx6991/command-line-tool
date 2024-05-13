// deprecated: because go can not change caller working directory

package main

import (
	"log"
	"os"
)

func main() {
	dirName, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(os.Args) > 1 {
		dirName = os.Args[1]
	}

	err = os.Chdir(dirName)
	if err != nil {
		log.Fatal(err)
		return
	}
}
