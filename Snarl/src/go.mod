module github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game

go 1.15

replace (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level v0.0.0-20210217010842-114e00eb25a4 => ./level
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/render v0.0.0-20210217010842-114e00eb25a4 => ./render
)

require github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/render v0.0.0-20210217010842-114e00eb25a4 // indirect
