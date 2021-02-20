module github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests

go 1.15

require (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game v0.0.0
)

replace (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game v0.0.0 => ../src/Game
)