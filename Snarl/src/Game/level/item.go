package level

// Item ID Constants (exposed to be usasble)
const (
	NoItem       = iota
	KeyID        = iota
	LockedExit   = iota
	UnlockedExit = iota
)

type Item struct {
	Type int
}

func NewKey() Item {
	return Item{
		Type: KeyID,
	}
}
