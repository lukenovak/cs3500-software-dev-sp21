# MEMORANDUM

#### DATE: 2/2/2021
#### TO: Ferdinand Vesely, Project Manager
#### FROM: William Victoria and Luke Novak, Software Developers
#### SUBJECT: Issues Implementing Traveller Specification

---

We encountered multiple issues trying to implement this package specification, many of which would have made it impossible to return working code had we followed the spec to the letter. We attempted to write a package that delivered the functionality we believe the specification was attempting to dictate.

## Major Concerns

- only specifies interfaces to be implemented, no classes
  - > interfaces for `Character`, `Town`, and `Traveller` must be created
- no methods provided for Town interface
  - only specifies two fields
  - without methods in the town interface the problem is not solvable
- wants to add fields to interfaces, which makes them `public static final`, which limits their usefulness
  - > The interface should have a Boolean type field `is_occupied`
- the `world_graph` field in `Traveller` should have a type specified for the List, i.e. `Map<Town, List<Town>>`
  - > The interface should have a Map<Town, List<Town>> type field `world_graph`
- wants to add instance fields to interfaces, which is not possible
  - this spec specifically asked for interfaces, we tried to reconcile this by adding getters/setters to the interfaces and implementing classes for each interface


## Minor Concerns

- does not follow Java naming conventions
  - uses snake_case instead of camelCase
  - > The interface should have a `add_to_network` method
- avoid initializing fields to `null`
  - > This can be initialized as null.
- fields withing classes should be private, implementation details
- `place_character` should not be in the Traveller interface, as it does not relate to the Traveller, it should be in either the Town or Character interface
- `IllegalStateException` is not the correct exception to use in this case
