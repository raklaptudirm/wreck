// Copyright © 2022 Rak Laptudirm <rak@laptudirm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"laptudirm.com/x/wreck/pkg/board"
	"laptudirm.com/x/wreck/pkg/tablebase"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "usage: wreck [position]")
		os.Exit(1)
	}

	position := `.........`
	if len(os.Args) == 2 {
		position = os.Args[1]
	}

	b, err := board.New(position)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	table := tablebase.Generate()

	fmt.Println("The Wreck Tic-Tac-Toe Engine")
	fmt.Println("Copyright © 2022 Rak Laptudirm <rak@laptudirm.com>")
	fmt.Println("Licensed under the Apache License, Version 2.0")
	fmt.Println("\nType 'help' for help regarding commands")

	reader := bufio.NewReader(os.Stdin)
commands:
	for {
		fmt.Print("\nwreck :: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "wreck: error reading from stdin")
			os.Exit(1)
		}
		fmt.Println()

		args := strings.Split(strings.Trim(input, "\n\r\t "), " ")

		if len(args) == 0 {
			continue commands
		}

		switch args[0] {
		case "exit":
			break commands

		case "load":
			if len(args) != 2 {
				fmt.Println("wreck: usage: load <position>")
				break
			}

			b, err = board.New(args[1])
			if err != nil {
				fmt.Println(err)
				break
			}

			if index, found := table.IndexOf(b); found {
				fmt.Print(table.Get(index).String())
			} else {
				fmt.Println("wreck: current position not found in tablebase")
			}

		case "play":
			switch {
			case len(args) != 2:
				fmt.Println("wreck: usage: play <move>")
			case len(args[1]) != 1:
				fmt.Printf("wreck: %#v is not a valid move\n", args[1])
			default:
				move := args[1][0]
				err := b.Play(board.Move(move - 48))
				if err != nil {
					fmt.Println(err)
					break
				}

				if index, found := table.IndexOf(b); found {
					fmt.Print(table.Get(index).String())
				} else {
					fmt.Println("wreck: current position not found in tablebase")
				}
			}

		case "eval":
			if len(args) != 1 {
				fmt.Println("wreck: usage: eval")
				break
			}

			if index, found := table.IndexOf(b); found {
				fmt.Print(table.Get(index).String())
			} else {
				fmt.Println("wreck: current position not found in tablebase")
			}

		case "help":
			helpString := `Commands:
  load <position>   Load the given position into wreck
  play <move>       Play the given move on the current position
  eval              Evaluate the current position and show data
  exit              Exit from the repl

Position String (<position>):
  A position in wreck is represented by a 9-character string which is
  composed of the symbols x, o, and . which represent a mark by player x, a
  mark by player o, and an empty cell. Each character represents a cell in
  the tic tac toe board.

Moves (<move>):
  Moves are represented by the numbers 1-9 where each number represents a
  position in the tic tac toe board.
    1 2 3
    4 5 6
    7 8 9`
			fmt.Println(helpString)

		default:
			fmt.Printf("wreck: unknown command %#v\n", args[0])
		}
	}
}
