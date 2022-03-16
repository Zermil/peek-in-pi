package main

import (
	"fmt"
	"os"
	"strings"
)

func error(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
	os.Exit(1)
}

func main() {
	content, err := os.ReadFile("./pi.txt")

	if err != nil {
		error(err.Error())
	}

	if len(os.Args) == 1 {
		error("Insufficient number of arguments")
	}

	sequence := os.Args[1]
	str := string(content)
	
	if strings.Contains(str, sequence) {
		index := strings.Index(str, sequence)
		index_after := index + len(sequence)
		
		fmt.Printf("Found %s at: %d\n",
			sequence + str[index_after:index_after + 6] + "...",
			index + 1)
	} else {
		fmt.Printf("Not possible to find '%s'", sequence)
	}
}
