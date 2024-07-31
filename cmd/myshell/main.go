package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		s, err := bufio.NewReader(os.Stdout).ReadString('\n')

		if err != nil {
			fmt.Println("panic")
			os.Exit(1)
		}
		s = strings.TrimSpace(s)
		if strings.Contains(s, "exit 0") {
			os.Exit(0)
		}
		fmt.Fprint(os.Stdout, strings.TrimSpace(string(s))+": command not found\n")

	}
}
