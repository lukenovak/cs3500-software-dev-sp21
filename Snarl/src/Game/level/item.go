package level

// Item ID Constants (exposed to be usasble)
const (
	NoItem = 0
	KeyID = 1
)

type Item struct {
	Type int
}

func NewKey() Item {
	return Item{
		Type: KeyID,
	}
}