module cmd

go 1.15

require github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json v0.0.0
require github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/tester v0.0.0

replace github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json v0.0.0 => ../json
replace github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/tester v0.0.0 => ../tester

