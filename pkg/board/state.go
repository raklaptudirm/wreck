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

// State represents the state of a tic tac toe Board.
type State int

// Constants representing various states of a Board.
const (
	Unfinished = iota
	GameDrawn
	PlayerXWon
	PlayerOWon
)

// String converts a Board state into it's string representation.
func (s State) String() string {
	switch s {
	case Unfinished:
		return "unfinished game"
	case GameDrawn:
		return "draw"
	case PlayerXWon:
		return "x wins"
	case PlayerOWon:
		return "o wins"
	default:
		return "invalid board state"
	}
}
