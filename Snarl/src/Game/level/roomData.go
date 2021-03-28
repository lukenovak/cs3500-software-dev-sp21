package level

type RoomGraphNode interface {
	GetId() int                                    // returns the room or hallway ID. Useful for surveying tiles in a specific room
	ConnectNode(connectingNode RoomGraphNode)      // Adds the given room as a connection to the room it's called on, and vice versa
	insertConnection(connectingNode RoomGraphNode) // Inserts a single node into a room graph's connections. Internal Use Only!
	Type() string                                  // returns a string representing the type of the room
	GetConnections() []RoomGraphNode               // Gets the list of this room's connections
	GetStartPoint() Position2D
}

/* ------------------------- Room Metadata ------------------------------- */

// Keeps track of room metadata
type RoomData struct {
	Id          int
	TopLeft     Position2D
	Size        Position2D
	Connections []RoomGraphNode
}

func (room RoomData) GetId() int {
	return room.Id
}

func (room RoomData) ConnectNode(connectingNode RoomGraphNode) {
	room.Connections = append(room.Connections, connectingNode)
	connectingNode.insertConnection(room)

}

func (room RoomData) insertConnection(connectingNode RoomGraphNode) {
	room.Connections = append(room.Connections, connectingNode)
}

func (room RoomData) Type() string {
	return "room"
}

func (room RoomData) GetConnections() []RoomGraphNode {
	return room.Connections
}

func (room RoomData) GetStartPoint() Position2D {
	return room.TopLeft
}

/* ------------------------- Hallway Metadata ------------------------------- */

// Keeps track of hallway metadata, which is slightly different than room metadata
type HallData struct {
	Id          int
	Start       Position2D
	End         Position2D
	Waypoints   []Position2D
	Connections []RoomGraphNode
}

func (hall HallData) GetId() int {
	return hall.Id
}

func (hall HallData) ConnectNode(connectingNode RoomGraphNode) {
	hall.Connections = append(hall.Connections, connectingNode)
	connectingNode.insertConnection(hall)
}

func (hall HallData) insertConnection(connectingNode RoomGraphNode) {
	hall.Connections = append(hall.Connections, connectingNode)
}

func (hall HallData) Type() string {
	return "hallway"
}

func (hall HallData) GetConnections() []RoomGraphNode {
	return hall.Connections
}

func (hall HallData) GetStartPoint() Position2D {
	return hall.Start
}
