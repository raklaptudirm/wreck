# The Wreck Tic-Tac-Toe Engine
Wreck is a tic-tac-toe analysis engine which is capable of perfect play,
i.e, it can win any winnable, and defend any defendable position. A full
game played against Wreck will either end in a draw or a win for Wreck.

### Installation

``` bash
go install laptudirm.com/x/wreck/cmd/wreck@latest
```

### Usage

#### Main Command
```bash
wreck [position]
```

#### REPL Commands
```bash
wreck :: help            # help regarding commands and the repl
wreck :: load <position> # load this position into the engine
wreck :: play <move>     # play the provided move on the current position
wreck :: eval            # evaluate current position
wreck :: exit            # exit from program
```

### Evaluation
Wreck evaluates position as a number. An evaluation of `±00` means the
position is equal, and perfect play will result in a draw. An evaluation
starting with a `+`, like `+Wn` means  means player `X` will win in `n`
steps, and an evaluation starting with `-` means player `O` will win in `n`
steps.

### Position Strings
A tic tac toe position is represented by a 9-character long position string
which is composed of `x`, `o`, and `.` symbols. Each of the nine characters
represents one of the cells on a tic tac toe board, and the symbols represent
a mark by player `X`, a mark by player `O`, and an empty cell respectively.

```
Any tic-tac-toe position:
1 2 3  x o .
4 5 6  x . .
7 8 9  o . .

Is represented in the following format:
  123456789
  xo.x..o..
```

### Moves
A move on the tic tac toe board which is at a particular position is
represented by a number from 1-9, each of which represent a particular cell
on the board.
