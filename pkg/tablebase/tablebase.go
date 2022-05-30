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
	"sort"

	"laptudirm.com/x/wreck/pkg/board"
	"laptudirm.com/x/wreck/pkg/evaluation"
)

// tablebase is table of all possible tic tac toe board positions and their
// evaluations, i.e, how good they are for each player.
type tablebase struct {
	data [10][]boardData
}

func (t *tablebase) Search(b board.Board) (boardData, bool) {
	index, found := t.indexOf(b)
	if found {
		return index.fetch(), true
	}

	return boardData{}, false
}

// get fetches the BoardData present at the given boardIndex in the
// tablebase.
func (t *tablebase) get(index boardIndex) boardData {
	return t.data[index.move][index.index]
}

// indexOf fetches the boardIndex of a tic tac toe position from the
// tablebase. It returns false as the second argument if the position can't
// be found.
func (t *tablebase) indexOf(b board.Board) (boardIndex, bool) {
	move := b.MoveNumber()
	for i, data := range t.data[move] {
		if data.board == b {
			return boardIndex{
				move:  move,
				index: i,
				table: t,
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
		table: t,
	}
}

// boardIndex represents the index of a position in the tablebase.
type boardIndex struct {
	move  int // move number
	index int // tablebase index

	table *tablebase // parent tablebase
}

// fetch gets the boardData at the current index in the parent tablebase,
// and returns it.
func (i boardIndex) fetch() boardData {
	return i.table.get(i)
}

// boardData stores position metadata including the position itself and
// it's evaluation. It forms the nodes of the tablebase.
type boardData struct {
	board board.Board // position

	eval    evaluation.Abs // position evaluation from children
	moveMap                // moves mapped to resulting positions

	table *tablebase // parent tablebase
}

// String converts a BoardData instance to it's string representation.
func (b boardData) String() string {
	s := fmt.Sprintf("%s\n", b.board)
	switch b.board.State() {
	case board.Unfinished:
		if b.board.XsTurn() {
			s += "[turn of player x]"
		} else {
			s += "[turn of player o]"
		}

		s += "\n\nLine     : Evaluation\n"

		moves := b.Moves()
		for _, data := range moves {
			nextEval := data.index.fetch().eval
			s += fmt.Sprintf("  Move %d : %s\n", data.move, nextEval)
		}

	case board.PlayerXWon:
		s += "\n(player x won)\n"
	case board.PlayerOWon:
		s += "\n(player o won)\n"
	case board.GameDrawn:
		s += "\n(game drawn)\n"
	}

	return s
}

// Position returns a Board representing the position of the boardData.
func (b boardData) Position() board.Board {
	return b.board
}

// MoveData returns a boardData representing the position after the given
// move is made on the current board.
func (b boardData) MoveData(move board.Move) (boardData, bool) {
	if data, found := b.Search(move); found {
		return data.index.fetch(), true
	}

	return boardData{}, false
}

// AbsEval returns the absolute evaluation of the Board that this boardData
// represents.
func (b boardData) AbsEval() evaluation.Abs {
	return b.eval
}

// RelEval returns the relative evaluation of the Board that this boardData
// represents.
func (b boardData) RelEval() evaluation.Rel {
	return evaluation.ToRel(b.eval, b.board)
}

// moveMap maps all valid moves in a position to their corresponding
// boardData.
type moveMap struct {
	boardMap []moveMapEntry
}

// Moves returns an array of moveMapEntries sorted according to their
// evaluation from best to worst. A move lower than another move may also
// have an equivalent evaluation.
func (m moveMap) Moves() []moveMapEntry {
	return m.boardMap
}

// Search looks for a moveMapEntry in the moveMap which represents the
// given move.
func (m *moveMap) Search(target board.Move) (moveMapEntry, bool) {
	for _, move := range m.boardMap {
		if move.move == target {
			return move, true
		}
	}

	return moveMapEntry{}, false
}

// add adds the given move with the given boardIndex to the moveMap.
func (m *moveMap) add(move board.Move, index boardIndex) {
	eval := evaluation.Flip(index.fetch().RelEval())
	m.boardMap = append(m.boardMap, moveMapEntry{move: move, index: index, eval: eval})
}

// finalize signals that no more elements will be added to the moveMap, and
// initiates sorting of the entries according to their evaluation.
func (m *moveMap) finalize() {
	sort.Slice(m.boardMap, func(i, j int) bool {
		return m.boardMap[i].eval > m.boardMap[j].eval
	})
}

// moveMapEntry represents individual entries in a moveMap
type moveMapEntry struct {
	move  board.Move     // represented move
	index boardIndex     // board state after move
	eval  evaluation.Rel // move evaluation
}
