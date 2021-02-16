# MEMORANDUM

#### DATE: February 14, 2021
#### TO: Growl, Inc.
#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.
#### SUBJECT: Snarl Data Representations

---

This memo presents our data representations that will be used in our implementation 
of Snarl. We used the provided Snarl documentation including the overview, plan, 
and protocol documents to determine our data representations.

Our Game Manager stores the level data as a 2d array of tile pointers (`[][]*Tile`). 
This array is generated using 

- tiles
- objects
- players
  - player data
    - id
    - type (player, adversary)
    - hp
    - atk
  - location


- pickup object
- unlock exit
- move character
- generate character partial view
- get tile data
- get player info
- expel characters
- generate new game state
