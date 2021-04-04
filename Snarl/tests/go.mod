module github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests

go 1.15

require (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game v0.0.0
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807 // indirect
)

replace (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game v0.0.0 => ../src/Game/
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests v0.0.0 => ./
)
