package level

// RoomGraphNode represents a single node in the graph of RoomData and HallData that makes up a level's metadata graph
type RoomGraphNode interface {
	// GetId returns the room or hallway ID. Useful for surveying tiles in a specific room
	GetId() int

	// ConnectNode the given room as a connection to the room it's called on, and vice versa
	ConnectNode(connectingNode RoomGraphNode)

	// insertConnection inserts a single node into a room graph's connections. Internal Use Only!
	insertConnection(connectingNode RoomGraphNode)

	// Type returns a string representing the type of the room
	Type() string

	// GetConnections the list of this room's connections
	GetConnections() []RoomGraphNode

	// GetStartPoint returns the start point of this room or hallway
	GetStartPoint() Position2D
}

/* ------------------------- Room Metadata ------------------------------- */

// RoomData represents room metadata
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

// HallData contains hallway metadata, which is slightly different than room metadata
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
