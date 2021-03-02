package level

type RoomGraphNode interface {
	GetId() int	// returns the room or hallway ID. Useful for surveying tiles in a specific room
	ConnectNode(connectingNode RoomGraphNode) // Adds the given room as a connection to the room it's called on, and vice versa
	insertConnection(connectingNode RoomGraphNode) // Inserts a single node into a room graph's connections. Internal Use Only!
}

/* ------------------------- Room Metadata ------------------------------- */

// Keeps track of room metadata
type RoomData struct {
	Id			int
	Type 		string
	TopLeft 	Position2D
	Size		Position2D
	Connections []RoomGraphNode
}

func (room *RoomData) GetId() int {
	return room.Id
}

func (room *RoomData) ConnectNode(connectingNode RoomGraphNode) {
	room.Connections = append(room.Connections, connectingNode)
	connectingNode.insertConnection(room)

}

func (room *RoomData) insertConnection(connectingNode RoomGraphNode) {
	room.Connections = append(room.Connections, connectingNode)
}

/* ------------------------- Hallway Metadata ------------------------------- */

// Keeps track of hallway metadata, which is slightly different than room metadata
type HallData struct {
	Id			int
	Type 		string
	Start		Position2D
	End			Position2D
	Waypoints	[]Position2D
	Connections []RoomGraphNode
}

func (hall *HallData) GetId() int {
	return hall.Id
}

func (hall *HallData) ConnectNode(connectingNode RoomGraphNode) {
	hall.Connections = append(hall.Connections, connectingNode)
	connectingNode.insertConnection(hall)
}

func (hall *HallData) insertConnection(connectingNode RoomGraphNode) {
	hall.Connections = append(hall.Connections, connectingNode)
}
