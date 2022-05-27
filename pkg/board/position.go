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

package board

import (
	"fmt"
	"strings"
)

// IsValidPosition verifies whether the given string is a valid tic tac toe
// position string. Note that this is just a simple check, and it
// classifies positions with multiple winners as valid. The final
// verification is whether the position is present in the tablebase or not.
func IsValidPosition(pos string) bool {
	// the position string's length should be 9, and it should be
	// entirely composed of x, o, and .
	if len(pos) != 9 || len(strings.Trim(pos, "xo.")) != 0 {
		return false
	}

	xCount := strings.Count(pos, "x")
	oCount := strings.Count(pos, "o")

	// the number of xs should be equal to (x's turn) or one more than
	// (o's turn) the number of os.
	if xCount != oCount && xCount-1 != oCount {
		return false
	}

	return true
}

// PositionError is the error reported when an invalid tic tac toe position
// string is provided to some methods.
type PositionError struct {
	posString string
}

func (e PositionError) Error() string {
	return fmt.Sprintf("board: invalid position string %#v", e.posString)
}

// New creates a new Board with the given position. It returns a
// PositionError if the given position string is invalid.
//
// A tic tac toe position string is a string of length 9, where each
// character represents a cell on the board. The symbols x, o, and .
// represent a mark by player x, a mark by player o, and an empty cell
// respectively.
func New(pos string) (Board, error) {
	if !IsValidPosition(pos) {
		return Board{}, PositionError{pos}
	}

	var x Bitboard
	var o Bitboard

	var moves int

	// put marks on the bitboards
	for i, mark := range pos {
		move := Move(i + 1)

		switch mark {
		case 'x':
			x.Set(move)
			moves++
		case 'o':
			o.Set(move)
			moves++
		case '.':
		default:
			// unreachable
			panic("board: invalid position after verification")
		}
	}

	b := Board{
		x: x,
		o: o,

		moveNum: moves,
	}

	// update state of board
	b.updateState()

	return b, nil
}
