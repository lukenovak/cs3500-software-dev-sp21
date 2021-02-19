module github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/tester

go 1.15

require (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level v0.0.0-20210219013218-e92ff2d14410
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json v0.0.0-20210219015015-a6f0bd0c508c
)

replace (
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json v0.0.0-20210219015015-a6f0bd0c508c => ../json
	github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level v0.0.0-20210219013218-e92ff2d14410 => ../../../src/level
)
