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

import "laptudirm.com/x/wreck/pkg/board"

// tablebase is table of all possible tic tac toe board positions and their
// evaluations, i.e, how good they are for each player.
type tablebase struct {
	data [10][]boardData
}

// boardIndex represents the index of a position in the tablebase.
type boardIndex struct {
	move  int // move number
	index int // tablebase index
}

// Get fetches the BoardData present at the given boardIndex in the
// tablebase.
func (t *tablebase) Get(index boardIndex) boardData {
	return t.data[index.move][index.index]
}

// IndexOf fetches the boardIndex of a tic tac toe position from the
// tablebase. It returns false as the second argument if the position can't
// be found.
func (t *tablebase) IndexOf(b board.Board) (boardIndex, bool) {
	move := b.MoveNumber()
	for i, data := range t.data[move] {
		if data.board == b {
			return boardIndex{
				move:  move,
				index: i,
			}, true
		}
	}

	// position not found
	return boardIndex{}, false
}

// pushBoard adds a BoardData entry to the tablebase.
func (t *tablebase) pushBoard(b boardData) boardIndex {
	move := b.board.MoveNumber()

	// add to tablebase
	t.data[move] = append(t.data[move], b)
	return boardIndex{
		move:  move,
		index: len(t.data[move]) - 1,
	}
}
