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

This update message is very similar to the `(player-update-message)` and differs only in the `level` field. This change
was required because adversaries receive the entire level, not just a localized layout. A `(level)` is defined in
Milestone 4.