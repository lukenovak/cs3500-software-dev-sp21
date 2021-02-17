module "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/cmd"

go 1.15

require (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level" v0.0.0
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/render" v0.0.0
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/item" v0.0.0

)

replace (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level" v0.0.0 => ../level
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/render" v0.0.0 => ../render
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/item" v0.0.0 => ../item
)