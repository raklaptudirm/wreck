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

// Package evaluation implements various functions and types relating to
// evaluating a tic tac toe position, and manipulating that evaluation. It
// also contains constants representing various evaluations.
package evaluation

import (
	"fmt"

	"laptudirm.com/x/wreck/pkg/board"
)

// eval represents the underlying type of all evaluation types.
type eval int8

// Rel represents a relative position evaluation.
type Rel eval

// relative evaluations representing various states
const (
	Draw    Rel = 0
	WinIn1  Rel = 10
	LossIn1 Rel = -10
)

// Abs represents an absolute position evaluation.
type Abs eval

// String returns the string representation of the given absolute
// evaluation.
func (a Abs) String() string {
	switch {
	case a == 0:
		return "±00"
	case a > 0:
		steps := 11 - a
		return fmt.Sprintf("+W%d", steps)
	case a < 0:
		steps := 11 + a
		return fmt.Sprintf("-W%d", steps)
	default:
		return "invalid"
	}
}

// ToRel converts a absolute position evaluation, where positive numbers
// represent a win for x and negative numbers represent a win for o to a
// turn relative evaluation where a positive number represents a win and a
// negative number represents a loss for the current player.
func ToRel(a Abs, b board.Board) Rel {
	r := Rel(a)
	if b.XsTurn() {
		return r
	}

	return -r
}

// ToAbs converts a turn relative evaluation where a positive number
// represents a win and a negative number represents a loss for the current
// player to an absolute position evaluation, where positive numbers
// represent a win for x and negative numbers represent a win for o.
func ToAbs(r Rel, b board.Board) Abs {
	a := Abs(r)
	if b.XsTurn() {
		return a
	}

	return -a
}

// Flip flips a relative evaluation to be from the perspective of the
// opponent, where a win for the current player turns into a loss for the
// opponent and vice-versa.
func Flip(e Rel) Rel {
	if e = -e; e > Draw {
		e--
	}

	return e
}
