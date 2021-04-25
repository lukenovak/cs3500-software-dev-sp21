# MEMORANDUM

#### DATE: April 23, 2021

#### TO: Growl, Inc.

#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.

#### SUBJECT: Snarl Remote Adversary Client

---

This memo presents implementation of the remote adversary client.

Our server registers adversary clients after registering player clients. A similar flow is used to register the
adversary clients. An adversary client connects to the server, and the server sends a `(server-welcome)` message. The
server then sends the string "type" to the client and the client sends a `(type)` to the server in response. Adversary
clients are assigned a random name in the format `"remote-{rand.Int()}"`. The server uses the same timeout for all
clients.

During the game the adversary clients behave similarly to the player clients. The primary difference is that the
adversary clients are sent an `(adversary-update-message)` instead of a `(player-update-message)`. The adversaries send
and receive moves in the same way as the player clients, but without user input for move selection. The adversaries all
move in order after all players have made a move.

### New JSON Message Definitions

`type`

A `(type)` is a JSON string. It is one of:

- "zombie"
- "ghost"

`adversary-update-message`

An `(adversary-update-message)` is a JSON object containing an adversary update and an optional message from the server.

```json
{
  "type": "adversary-update",
  "level": (level),
  "position": (point),
  "objects": (object-list),
  "actors": (actor-position-list),
  "message": (maybe-string)
}
```

By implementing a new type of update, the previous protocol is kept intact, while the new functionality is able to get
the information it needs.

This update message is very similar to the `(player-update-message)` and differs only in the `level` field. This change
was required because adversaries receive the entire level, not just a localized layout. A `(level)` is defined in
the testing task for Milestone 4, and works using a system of Rooms, Hallways, and layouts for each room.

### Backward Compatibility in running the executable

Finally, the executable is 100% backward compatible. This is accomplished by having the number
of remote adversaries set at the command line with a new flag. The game manager knows how many
adversaries are connected, as it has been extended with an `outsideAdversaries` field, and it
generates local adversaries to fill out any remaining spots not filled by the remote adversaries.
In the case that no remote adversaries are specified, the game manager would handle all of the
adversary behavior exactly as it did before.
