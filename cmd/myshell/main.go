package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
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
	{"pwd"},
	{"cd"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		s, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("panic")
			os.Exit(1)
		}
		s = strings.TrimSpace(s)

		// Komut ve argümanları ayırma
		parts := strings.Fields(s)
		if len(parts) == 0 {
			continue
		}

		cmdName := parts[0]
		args := parts[1:]

		if strings.HasPrefix(s, "type ") {
			if len(parts) < 2 {
				fmt.Fprint(os.Stdout, "type: missing operand\n")
				continue
			}
			cmd := strings.TrimSpace(parts[1])
			found := false

			for _, v := range commands {
				if v.name == cmd {
					fmt.Fprint(os.Stdout, cmd+" is a shell builtin\n")
					found = true
					break
				}
			}

			if !found {
				path := strings.Split(os.Getenv("PATH"), ":")
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

		if cmdName == "exit" && len(args) > 0 && args[0] == "0" {
			os.Exit(0)
		}

		if cmdName == "echo" {
			message := strings.Join(args, " ")
			fmt.Fprint(os.Stdout, message+"\n")
			continue
		}

		if cmdName == "pwd" {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error printing directory: %s\n", err)
				continue
			}
			fmt.Fprintf(os.Stdout, "%s\n", dir)
			continue
		}

		if cmdName == "cd" {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "cd: missing argument\n")
				continue
			}

			newDir := args[0]

			// Handle home directory shortcut (~)
			if newDir == "~" {
				homeDir := os.Getenv("HOME")
				if homeDir == "" {
					fmt.Fprintf(os.Stderr, "cd: HOME environment variable not set\n")
					continue
				}
				newDir = homeDir
			} else if len(newDir) > 1 && newDir[0] == '~' {
				// Handle the case where ~ is followed by a path (e.g., ~/dir)
				homeDir := os.Getenv("HOME")
				if homeDir == "" {
					fmt.Fprintf(os.Stderr, "cd: HOME environment variable not set\n")
					continue
				}
				newDir = path.Join(homeDir, newDir[1:])
			}

			err := os.Chdir(newDir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", newDir)
				continue
			}

			continue
		}

		// Dosya yolunu kontrol et
		path, err := exec.LookPath(cmdName)
		if err != nil {
			fmt.Fprint(os.Stdout, cmdName+": command not found\n")
			continue
		}

		// Komutu oluştur ve çalıştır
		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Fprint(os.Stdout, cmdName+": "+err.Error()+"\n")
		}
	}
}
