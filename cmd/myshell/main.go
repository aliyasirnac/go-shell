package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Command struct {
	name string
}

var commands = []Command{
	{"type"},
	{"exit"},
	{"echo"},
}

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		s, err := bufio.NewReader(os.Stdout).ReadString('\n')

		if err != nil {
			fmt.Println("panic")
			os.Exit(1)
		}
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "type ") { // type command takes an argument so we need to check string with HasPrefix command
			parts := strings.Split(s, " ")

			cmd := strings.TrimSpace(parts[1])
			found := false

			path := strings.Split(os.Getenv("PATH"), ":")

			for _, v := range commands {
				if strings.Contains(v.name, cmd) {
					fmt.Fprint(os.Stdout, cmd+" is a shell builtin\n")
					found = true
					break
				}
			}

			if !found {
				for _, dir := range path {
					cmdPath := filepath.Join(dir, cmd)
					if _, err := os.Stat(cmdPath); err == nil {
						fmt.Fprint(os.Stdout, cmd+" is "+cmdPath+"\n")
						found = true
						break
					}
				}
			}

			if !found {
				fmt.Fprint(os.Stdout, cmd+": not found\n")
			}
			continue
		}

		if strings.Contains(s, "exit 0") {
			os.Exit(0)
		}

		if strings.HasPrefix(s, "echo") {
			message := strings.Split(s, "echo ")

			fmt.Fprint(os.Stdin, strings.Join(message, "")+"\n")
			continue
		}

		fmt.Fprint(os.Stdout, s+": command not found\n")
	}
}
