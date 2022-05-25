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

// BitBoard represents a tic tac toe board where each cell can be one of
// two states, set or not set.
type Bitboard struct {
	uint16
}

// NewBitboard creates a new Bitboard from the given state.
func NewBitboard(state uint16) Bitboard {
	return Bitboard{state}
}

// String converts a Bitboard to it's string representation.
func (b Bitboard) String() string {
	var s string
	for i := Move(1); i <= 9; i++ {
		if b.Has(i) {
			s += "x"
		} else {
			s += "."
		}
	}

	return s
}

// Has checks if the given position is set in the Bitboard.
func (b *Bitboard) Has(pos Move) bool {
	return b.uint16>>getPos(pos)&1 == 1
}

// Set sets the given position in the Bitboard.
func (b *Bitboard) Set(pos Move) {
	b.uint16 |= buffer(pos)
}

// Unset clears the given position in the Bitboard.
func (b *Bitboard) Unset(pos Move) {
	b.uint16 &^= buffer(pos)
}

// HasWon checks if one of the given lines are completely set in the
// Bitboard. In a player's bitboard, it checks if the player has won.
func (b *Bitboard) HasWon() bool {
	// ways in which a player can win
	winningLines := [][3]Move{
		// rows
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},

		// columns
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},

		// diagonals
		{1, 5, 9},
		{3, 5, 7},
	}

checkingForWins:
	for _, line := range winningLines {
		// check if all three positions are set
		for i := 0; i < 3; i++ {
			if !b.Has(line[i]) {
				continue checkingForWins
			}
		}

		return true
	}

	return false
}

// buffer converts a given move into a flag buffer. A flag buffer is a
// buffer with some target bits set. Here, it is the position.
func buffer(pos Move) uint16 {
	return 1 << getPos(pos)
}

// getPos converts a Move into a position on the Bitboard.
func getPos(pos Move) uint8 {
	return 9 - uint8(pos)
}
