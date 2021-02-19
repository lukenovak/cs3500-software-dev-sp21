module github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/render

go 1.15

replace (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/item v0.0.0 => ../item
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level v0.0.0 => ../level
)

require github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level v0.0.0-20210217010842-114e00eb25a4 // indirect
