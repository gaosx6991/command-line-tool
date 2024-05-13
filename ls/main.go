// ls command
// example:
//	ls
// 	ls /usr/local

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dirName := "."

	if len(os.Args) > 1 {
		dirName = os.Args[1]
	}

	files, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
