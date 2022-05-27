# The Wreck Tic-Tac-Toe Engine
Wreck is a tic-tac-toe analysis engine which is capable of perfect play,
i.e, it can win any winnable, and defend any defendable position. A full
game played against Wreck will either end in a draw or a win for Wreck.

### Installation

``` bash
git clone https://github.com/raklaptudirm/wreck.git
cd wreck
go build ./cmd/wreck
./wreck # put this executable in your path
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
Wreck evaluates position as a number. An evaluation of `0` means the
position is equal, and perfect play will result in a draw. A positive number
means the position is winning for `x`, and perfect play will result in a win
for `x`, and vice-versa for negative numbers.

### Position Strings
A tic tac toe position is represented by a 9-character long position string
which is composed of `x`, `o`, and `.` symbols. Each of the nine characters
represents one of the cells on a tic tac toe board, and the symbols represent
a mark by player X, a mark by player O, and an empty cell respectively.

```
The following position:
x o .  1 2 3
x . .  4 5 6
o . .  7 8 9
Is represented by the following string:
xo.x..o..
123456789
```

### Moves
A move on the tic tac toe board which is at a particular position is
represented by a number from 1-9, each of which represent a particular cell
on the board.
