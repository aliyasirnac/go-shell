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

		if strings.Contains(s, "echo") {
			message := strings.Split(s, "echo ")

			fmt.Fprint(os.Stdin, strings.Join(message, "")+"\n")
			continue
		}

		fmt.Fprint(os.Stdout, s+": command not found\n")
	}
}
