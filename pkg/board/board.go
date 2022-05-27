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

import "fmt"

// Board represents the state of a tic tac toe board at any given point in
// time. It also stores additional metadata about the position.
// The zero value is a valid and usable Board.
type Board struct {
	// position
	x Bitboard
	o Bitboard

	// metadata
	moveNum int   // current move number
	state   State // state of the game
}

// Move represents a move on the Board. The numbers 1-9 represent the 9
// different playable positions on the board.
type Move uint8

// String converts a Board to it's string representation.
func (b Board) String() string {
	var s string
	for i := Move(1); i <= 9; i++ {
		// convert current cell to a symbol
		var symbol string
		switch {
		case b.x.Has(i):
			symbol = "x"
		case b.o.Has(i):
			symbol = "o"
		default:
			symbol = "."
		}

		s += symbol
		if i%3 == 0 {
			// separate by newline at row end
			if i != 9 {
				s += "\n"
			}
		} else {
			// separate by space otherwise
			s += " "
		}
	}

	return s
}

// InvalidMove represents an invalid move provided to Play.
type InvalidMove struct {
	move Move
}

func (e InvalidMove) Error() string {
	return fmt.Sprintf("play: invalid move %d", e.move)
}

// Play makes the given move on it's Board, and updates the position and
// state accordingly.
func (b *Board) Play(move Move) error {
	// check if move is valid
	if !b.IsValidMove(move) {
		return InvalidMove{move}
	}

	if b.XsTurn() {
		b.x.Set(move)
	} else {
		b.o.Set(move)
	}

	// increase move count
	b.moveNum++

	b.updateState()
	return nil
}

// updateState checks for wins or draws in the Board and updates the state
// accordingly.
func (b *Board) updateState() {
	switch {
	case b.x.HasWon():
		// x won
		b.state = PlayerXWon
	case b.o.HasWon():
		// o won
		b.state = PlayerOWon
	case b.moveNum == 9:
		// all moves completed without anyone winning
		// therefore position is a draw
		b.state = GameDrawn
	default:
		// game still ongoing
		b.state = Unfinished
	}
}

// ValidMoves calculates the valid moves in current position and returns
// them as a slice of Moves.
func (b *Board) ValidMoves() []Move {
	var moves []Move

	// game must be unfinished for there to be valid moves
	if b.state == Unfinished {
		for i := Move(1); i <= 9; i++ {
			if b.IsValidMove(i) {
				moves = append(moves, i)
			}
		}
	}

	return moves
}

// XsTurn checks whether it is x's turn to play and returns a bool
// accordingly.
func (b *Board) XsTurn() bool {
	return b.moveNum%2 == 0
}

// IsValidMove checks if the given move is valid on it's Board.
func (b *Board) IsValidMove(move Move) bool {
	switch {
	// check that game is unfinished, and that the move is on a valid and empty cell
	case b.state != Unfinished, move > 9 || move < 1, b.x.Has(move), b.o.Has(move):
		return false
	default:
		return true
	}
}

// MoveNumber returns the number of moves played on the Board.
func (b *Board) MoveNumber() int {
	return b.moveNum
}

// State returns the current state of the Board.
func (b *Board) State() State {
	return b.state
}
