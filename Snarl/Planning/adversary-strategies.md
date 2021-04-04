# Adversary Strategies

Our adversaries can be split into two main categories: Ghosts and Zombies

## Zombie

Zombies act "stupidly", which is to say that they move in random directions
unless a player is in the same room as them.

The logic works roughly as follows.

If a player is in the same room as a zombie: get all valid moves, and check them
to determine which one creates the minimum manhattan distance to any player. Make
that move.

If a player is not in the same room as a zombie: make a random valid move in any
direction.

## Ghost

A Ghost type adversary selects its moves in the following priority order:

1. If in the same room as a Player(s), move toward closest player. 
2. If adjacent to a Wall, teleports to a random wall tile
3. Otherwise, makes a random valid move.
