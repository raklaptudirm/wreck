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

import "laptudirm.com/x/wreck/pkg/board"

type Eval int8

// relative evaluations representing various states
const (
	Win  Eval = 1
	Draw Eval = 0
	Loss Eval = -1
)

// Reflect converts a absolute position evaluation, where positive
// numbers represent a win for x and negative numbers represent a win for o
// to a turn relative evaluation where a positive number represents a win
// and a negative number represents a loss for the current player, or vice-
// versa, depending on what type of evaluation was given to it.
func Reflect(e Eval, b board.Board) Eval {
	if b.XsTurn() {
		return e
	}

	return -e
}

// Flip flips a relative evaluation to be from the perspective of the
// opponent, where a win for the current player turns into a loss for the
// opponent and vice-versa.
func Flip(e Eval) Eval {
	return -e
}
