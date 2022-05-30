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
	"laptudirm.com/x/wreck/pkg/board"
	"laptudirm.com/x/wreck/pkg/evaluation"
)

// Generate creates and evaluates all the boards from the default tic tac
// toe starting position. It generates the entire tablebase.
func Generate() tablebase {
	var table tablebase
	var board board.Board // zero value is starting board

	// generate boards from starting position
	table.generateBoardsFrom(board)
	return table
}

// generateBoardsFrom generates the children Boards for a given Board,
// which consist of positions after playing a valid move. This is done
// recursively generating all the possible ways the game could progress
// from the given position. The generated boards are given an evaluation
// and stored in the tablebase. It returns the index of the given board in
// the tablebase and boards evaluation relative to the player.
func (t *tablebase) generateBoardsFrom(b board.Board) (boardIndex, evaluation.Rel) {
	// check if Board has already been generated
	if index, found := t.indexOf(b); found {
		eval := evaluation.ToRel(index.fetch().eval, b)
		return index, eval
	}

	var moves moveMap          // map of valid moves to boards
	eval := evaluation.LossIn1 // uses lowest possible eval

	// generate evaluation from game state
	switch b.State() {
	case board.Unfinished:
		validMoves := b.ValidMoves()

		// generate and evaluate each valid move on the Board
		for _, move := range validMoves {
			newBoard := b       // create a copy
			newBoard.Play(move) // play the move

			nextIndex, nextEval := t.generateBoardsFrom(newBoard)
			moves.add(move, nextIndex) // add move to moveMap

			// flip the relative eval to current player's perspective
			moveEval := evaluation.Flip(nextEval)

			// update board evaluation
			if moveEval > eval {
				eval = moveEval
			}
		}

	// game finished, no valid moves remain
	// so hardcode evaluation depending on state
	case board.PlayerXWon, board.PlayerOWon:
		// relative evaluation of an immediate loss
		eval = evaluation.LossIn1
	case board.GameDrawn:
		eval = evaluation.Draw
	}

	// finalize move map by sorting according to eval
	moves.finalize()

	// push given board to tablebase
	return t.pushBoard(boardData{
		board:   b,
		eval:    evaluation.ToAbs(eval, b),
		moveMap: moves,
		table:   t,
	}), eval
}
