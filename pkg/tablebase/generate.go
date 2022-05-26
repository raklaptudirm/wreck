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

package tablebase

import (
	"fmt"

	"laptudirm.com/x/wreck/pkg/board"
)

// Generate creates and evaluates all the boards from the default tic tac
// toe starting position. It generates the entire tablebase.
func Generate() tablebase {
	var table tablebase
	table.generateAllBoards()
	return table
}

// boardData stores position metadata including the position itself and
// it's evaluation. It forms the nodes of the tablebase.
type boardData struct {
	board board.Board // position

	eval  evaluation                // position evaluation from children
	moves map[board.Move]boardIndex // moves mapped to resulting positions

	table *tablebase // parent tablebase
}

// String converts a BoardData instance to it's string representation.
func (b boardData) String() string {
	s := fmt.Sprintf("%s\n\nEvaluation: %d\n\nLines:\n", b.board, b.eval)

	for move, boardIndex := range b.moves {
		nextEval := b.table.Get(boardIndex).eval
		s += fmt.Sprintf("%d: %d\n", move, nextEval)
	}

	return s
}

// Position returns a Board representing the position of the boardData.
func (b boardData) Position() board.Board {
	return b.board
}

// Evaluation returns the position evaluation of the boardData.
func (b boardData) Evaluation() evaluation {
	return b.eval
}

// MoveData returns a boardData representing the position after the given
// move is made on the current board.
func (b boardData) MoveData(move board.Move) (boardData, bool) {
	if index, found := b.moves[move]; found {
		return b.table.Get(index), true
	}

	return boardData{}, false
}

// evaluation represents the eval metric of a Board in the tablebase.
type evaluation = int8

// generateAllBoards calls the generateBoardsFrom function with the default
// Board, generating the entire tablebase.
func (t *tablebase) generateAllBoards() {
	var board board.Board // starting board
	t.generateBoardsFrom(board)
}

// generateBoardsFrom generates the children Boards for a given Board,
// which consist of positions after playing a valid move. This is done
// recursively generating all the possible ways the game could progress
// from the given position. The generated boards are given an evaluation
// and stored in the tablebase. It returns the index of the given board in
// the tablebase and it's evaluation.
func (t *tablebase) generateBoardsFrom(b board.Board) (index boardIndex, eval evaluation) {
	// check if Board has already been generated
	if index, found := t.IndexOf(b); found {
		return index, t.Get(index).eval
	}

	// map of valid moves to their resultant Board
	moveMap := make(map[board.Move]boardIndex)

	// generate evaluation from game state
	switch b.State() {
	case board.Unfinished:
		validMoves := b.ValidMoves()
		xsTurn := b.XsTurn()

		// start with lowest possible eval
		if xsTurn {
			// -1 is winning for o
			eval = -1
		} else {
			// 1 is winning for x
			eval = 1
		}

		// generate and evaluate each valid move on the Board
		for _, move := range validMoves {
			newBoard := b       // create a copy
			newBoard.Play(move) // play the move

			currIndex, currEval := t.generateBoardsFrom(newBoard)
			moveMap[move] = currIndex

			// update board evaluation
			if xsTurn {
				// x's turn, so use highest eval
				if currEval > eval {
					eval = currEval
				}
			} else {
				// o's turn, so use lowest eval
				if currEval < eval {
					eval = currEval
				}
			}
		}

	// game finished, no valid moves remain
	// so hardcode evaluation depending on state
	case board.PlayerXWon:
		eval = 1
	case board.PlayerOWon:
		eval = -1
	case board.GameDrawn:
		eval = 0
	}

	// push given board to tablebase
	index = t.pushBoard(boardData{
		board: b,
		eval:  eval,
		moves: moveMap,
		table: t,
	})

	return
}
