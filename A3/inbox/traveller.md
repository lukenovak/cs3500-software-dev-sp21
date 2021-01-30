### Traveller Assignment
Using Java 1.8.0_262, a route planner must be implemented for our client. Within a new package named `pathfinding`, interfaces for `Character`, `Town`, and `Traveller` must be created. The high-level concept, is that a simple-graph of `Towns` holds an unspecified number of `Character` implementations. Using an implementation of `Traveller`, our client must be able to tell if a given `Character` is capable of reaching a specified `Town` without being blocked by another `Character`.

#### Character Interface
A character represents a player in the game.
* The interface should have a field `current_location`, which represents what `Town` it currently exists in. This can be initialized as null. Include getter and setter methods.

#### Town Interface
A town represents a location in the game. Additionally, it represents a *node* in the game's world.
* The interface should have a Boolean type field `is_occupied`, which tells us if a `Character` occupies the `Town` or not. This can be initialized as *false*.
* The interface should have a String type field `name`, which represents the name of the `Town`. This can be initialized as null.

#### Traveller Interface
A *traveller* represents a graph handler which is capable of holding the game's world, and determining if a `Character` can travel the world without interference.
* The interface should have a Map<Town, List<Town>> type field `world_graph`, which represents each `Town` in the world, and what `Towns` are connected to it. This field should be initialized as an empty map object (an *empty graph*).
* The interface should have a `add_to_network` method, which has a **void** return type, and accepts a **Town**, as well as a list of **Towns** which is connected to.
  * This method is responsible for adding the `Town`, and it's connected *nodes* to the map which represents the game world's graph.
  * Ignore any call to this function which attempts to add a `Town` that already exists in the map.
* The interface should have a `place_character` method, which has a **Void** return type, and accepts a **Character** object, as well as a **Town** object.
  * This method is responsible for adding the Character to the provided town.
  * In the case that a Character is attempted to add to a town whose `is_occupied` field is *true*, an *IllegalStateException* should be thrown.
* The interface should have a `can_travel_to` method, which has a **Boolean** return type, and accepts a **Character** object, as well as a **Town** object.
  * This method is responsible for determining if the `Character` can travel from it's starting location to the provided `Town`, without needing to pass through any occupied towns.
  * In the case that a Character can travel from it's starting location to the desired `Town`, return *true*.
  * In the case that the Character cannot travel to the given `Town` in a valid manner, return *false*.
