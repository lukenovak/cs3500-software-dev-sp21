# MEMORANDUM

#### DATE: February 8, 2021
#### TO: Growl, Inc.
#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.
#### SUBJECT: Snarl Project Analysis

Writing conditions:
Design Task

>Read the description of Snarl, a custom designed game.

done.

>We will develop the software to run games of Snarl and tournaments for automated players adversaries. In a tournament, other people will supply the code for adversaries, and you will provide the framework for running games for these automated entities. The goal is to run tournaments where every signed up player gets a chance to participate in several Snarl games against these new enemies, in a manner yet to be determined.

>Write up a project analysis. The analysis should consist of two parts:

    Part 1 describes the identifiable components of your software system. Ask yourself

        what are the pieces that make up a player

        what are the pieces that make up an automated adversary

        what are the pieces that make up the game software

        “who” knows what, “who” needs to know what, and how do “they” communicate

        For the communication part, sketch a UML-ish sequence diagram in local.md.

        what common knowledge is necessary to make communication successful.

>Note: This prose uses “who” and “they” in reference to pieces of software. Think of them as possibly independent “actors” and humanize them temporarily. This is standard practice in techincal writing.

Based on your description of Snarl, we have drafted a high-level design of the 
game and the data it will represent.

Beginning at the highest level, a data structure we will refer to as a `Game`
will hold a few pieces of data: a `Dungeon`, that represents all of the map
data, and a list of "actors", which include the player characters and the adversary
monsters. These "actors" will be called `Mobs`. A `Game` will also have a
system for moving everything and advancing the turn. For the Player characters, 
the move instructions will be given by the the human player, but for
adversaries, these instructions will be determined by a `Conductor` that will
have access to the `Game` data to be able to determine how the adversaries
should move. The conductor, in theory, knows everything about the game, including
dungeon layout, player position, and enemy positions and will serve as the main 
processing unit for the enemy characters.

`Mobs` can be under one of two categories: user-controlled and computer-controlled.
Both groups share many characteristics, which is why Player characters are also
considered `Mobs`. However, given that their movement is dictated by the player (and
that there might be additional characteristics of a player character that adversaries
will never have) we have separated players out into a different category. All `Mobs`
have the potential to represent a great diversity of data, and the modular nature
of the design allows for new features to be added to players, adversaries, or both.
Some potential features this system can support include (but are not limited to):
- Inventory
- Health points
- Attack points
- Player/advsersary actions.

Importantly, `Mobs`, whether they are `Players` or non-player `ComputerMobs`, only
"know" their own position, and any data about themselves (like inventory or HP). This
data will be communicated back to the `Game` to update how the `Game` handles players.
For instance, when moving, a `Mob` can either be moved by the game directly, or ask the
`Game` if a move is legal. The `Mob` itself cannot ascertain the legality of moves
since it does not have access to information about the surrounding tiles.

`Dungeons` represent the game's map data, and are made up of a number of `Levels`. Each of
these `Levels` are made up of `Rooms`, which can be of a `Room` or `Hallway` type. `Rooms`
are made up of `Tiles`, which represent the base level of movement in Snarl. `Tiles` carry
a few pieces of data: The tile type, whether or not the tile is occupied, whether the `Tile`
is visible to the players and any items that may be on the tile. 
`Items` will be identified on `Tiles` by an ID number, which can
generate an `Item` in a `Player`s inventory. For now, the only item outlined in the spec is
the key, but we suspect that this functionality will be good to support future additions to
the game. Doors will be a special type of `Tile` that will check for the key in the player's
inventory, and if the player has it, they will unlock the door and proceed through. This action
will change the active `Level` until the last level of the `Dungeon` when the `Game` will note that
there are no more `Levels` in that Game's `Dungeon` and end the game, with the players winning.


Part 2:

In terms of development, we have a number of milestones to achieve. Each milestone will have a
"playable" demo prototype, showcasing the features added at that development step.

1. Map Generation and basic player movement
    - the Player character can spawn and move around within a map
2. Complete Dungeons with keys
    - The Player can navigate around 


    Part 2 describes how you should proceed about implementing these pieces. You will propose milestones for the project. Keep in mind that you wish to have “demo” software soon so that a potential client can admire fully working prototypes.

>The text of both Parts 1 and 2 (i.e. not the sequence diagram) must be in plan.md.

>Each memo file must not exceed one page.

Stuff to Represent:
- Conductor
- Mob
  - Sub-types
    - Player
    - Monsters
      - Ghost
      - Zombie
  - Data
    - Expulsion
    - Inventory
    - Health
    - Attack Points
    - Actions
- Zone
  - Hallways
  - Rooms
    - Tiles (items, actors)
    - Door (lockable)
    - Exit
- Game
- Dungeon (full map)
- Levels
- Items
  - Key
