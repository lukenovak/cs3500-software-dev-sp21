module github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/cmd

go 1.15

replace (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/item => ../item
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level => ../level
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/render => ../render
)

require (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/render v0.0.0-20210217010842-114e00eb25a4 // indirect
	fyne.io/fyne/v2 v2.0.0

)
