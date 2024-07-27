package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	s, err := bufio.NewReader(os.Stdout).ReadBytes('\n')

	if err != nil {
	}

	fmt.Fprint(os.Stdout, strings.TrimSpace(string(s))+": command not found\n")
}
