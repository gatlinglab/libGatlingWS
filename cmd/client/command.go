package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandRun() {
	reader := bufio.NewReader(os.Stdin)
	bQuit := false

	for {
		if bQuit {
			break
		}
		fmt.Print("> ")

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Remove the newline character.
		input = strings.TrimSuffix(input, "\n")

		// Skip an empty input.
		if input == "" {
			continue
		}

		switch input {
		case "quit", "exit":
			bQuit = true
		default:
			fmt.Println(g_socket)
			g_socket.WriteBinary([]byte(input))

		}
	}
}
