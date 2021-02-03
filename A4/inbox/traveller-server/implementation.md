# Implementation

## Structs

We made network `type Network []*Town` instead of `type Network []Town`.
This is to let us mutate towns in our functions if needed.

We also changed towns to use pointers of other structs:

```go
type Town struct {
  Name          string
  Character     *Character
  AdjacentTowns []*Town
}
```

Pointers make more sense because otherwise go will create copies of these fields, resulting in inconsistencies from mutation.

## Functions

The following are the new function signatures:

```go
// Creates a new Town struct in the given Network with the given name and returns a new network and a town pointer
func CreateTown(name string, n Network) Network

// Link two towns by adding each other to their AdjacentTowns lists
func LinkTowns(t1 *Town, t2 *Town)

// Unlinks two connected towns by removing each other from their AdjacentTowns lists
func UnlinkTowns(t1 *Town, t2 *Town)

// Creates a new Character with the given name in a town if the town does have a character
func PlaceNewCharacter(t *Town, name string) (*Character, error)

// Places an existing Character into a town if the town does not have a character
func PlaceExistingCharacter(t *Town, c *Character)

// Removes a Character from a given town, returning that character
func RemoveCharacterFromTown(t *Town) *Character

// This function determines if the given Character can travel to the given Town in the given Network
// Without passing through any other Town with a Character
func CanTravel(c *Character, t *Town, n Network) bool
```

- For `CreateTown`, it would return a new `Network` with the specified `Town` appended and the pointer to the new town. This made more sense rather than returning just the town.

- `LinkTowns` and `UnlinkTowns` implied mutation on the town parameters, so they required pointer parameters.

- `PlaceNewCharacter` implied mutation and the creation of an instance of a `Character` so the parameter is a town pointer and it returns a character pointer.

- `PlaceExistingCharacter` and `RemoveCharacterFromTown` implied mutation so the parameters were changed to pointers.

- `CanTravel` uses pointers as a way to efficiently use BFS on the network.

Note: Pointers are essentially required to use this because all the signatures from the traveller specs imply mutation. Without pointers go would pass in value copies of the parameters resulting in no mutation being done.
