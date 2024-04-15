//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	commands := 0

	for {
		cmd, err := r.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Exiting...")
			fmt.Printf("Commands: %v\n", commands)
			break
		}
		cmd = cmd[:len(cmd)-1]
		if cmd == "" {
			continue
		}

		commands++
		if cmd == "Q" || cmd == "q" {
			fmt.Println("Exiting...")
			fmt.Printf("Commands: %v\n", commands)
			break
		}

		if cmd == "hello" || cmd == "bye" {
			fmt.Println("Hello! or Bye! ;)")
		}

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
	}
}
