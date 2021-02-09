# MEMORANDUM

#### DATE: February 8, 2021
#### TO: Growl, Inc.
#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.
#### SUBJECT: Snarl Project Analysis

Based on your description of Snarl, we have drafted a high-level design of the 
game and the data it will represent.

Beginning at the highest level, a data structure we will refer to as a `Game`
will hold a few pieces of data: a `Dungeon`, that represents all of the map
data, and a list of "actors", which include the player characters and the adversary
monsters. These "actors" will be called `Mobs`. The `Game` will act as a central
system for moving everything and advancing the turn. For the Player characters, 
the move instructions will be given by the the human player, but for
adversaries, these instructions will be determined by a `Conductor` that will
have access to a `Game` state, including map data and character positions, determine it
should tell the `Game` to move the adversaries The conductor, serves as the main 
processing unit for the enemy characters, and a default local conductor will be
provided. This delegation of adversary movement to a `Conductor` will allow for
future extensibility to support tournaments with outside adversary control.

`Mobs` can be under one of two categories: user-controlled and computer-controlled.
Both groups share many characteristics, which is why Player characters are also
considered `Mobs`. However, given that their movement is dictated by the player (and
that there might be additional characteristics of a player character that adversaries
will never have) we have separated players out into a different category. `Mobs` right
now only carry their position data, but the modular design of the interface will allow
features to be added in the future. `Mobs` only "know" their own information, and nothing
about the world or tiles around them. That data is handled by the `Game` and the movement
of `Mobs` is carried out by the `Game` based on the position data of the `Mobs` and the
input from the player or the `Conductor`.

`Dungeons` represent the game's map data, and are made up of a number of `Levels`. Each of
these `Levels` are made up of `Rooms`, which can be of a `Room` or `Hallway` type. `Rooms`
are made up of `Tiles`, which represent the base level of movement in Snarl. `Tiles` carry
a few pieces of data: The tile type, whether or not the tile is occupied, whether the `Tile`
is visible to the players and any items that may be on the tile. For now, that would just be
the key item but could be expanded in the future. When a player character moves to the `key`,
the `Game` will notice the move and make the Locked Door `Tile` into an Unlocked Door `Tile`.

A rough order of operations as shown in our diagram works as follows:
1. The player starts the game, and a `Game` is created. The `Dungeon` and `Mobs` are generated
2. The player(s) gives their input to the game
3. The `Game` uses the input information to move the players or inform the player that
their move was illegal. If illegal, go back to step 2.  
3a. The player moves are *"resolved"* meaning that the `Game` checks/updates the `Dungeon` data.
4. The `Conductor` gets the game state from the `Game` and moves the adversaries  
3a. The adversaries' moves are "resolved"
5. If the game has not been won/lost, return to step two.

#### Developmemt Timeline:

In terms of development, we have a number of milestones to achieve. Each milestone will have a
"playable" demo prototype, showcasing the features added at that development step.

1. Room Generation and basic player movement
    - the Player character can spawn and move around within a Room (steps 1-3)
2. Complete Multi-Room Levels with keys
    - The Player can navigate around, finding the key, and exiting the level
3. Working enemies in levels (step 4)
4. Multi-level Dungeons
5. Remaining Polish
    - Menu
    - Additional Features.
