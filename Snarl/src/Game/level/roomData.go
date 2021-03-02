package level

type RoomData struct {
	Type string
	Connections []*RoomData
}

func (room *RoomData) InsertConnection(connectingRoom *RoomData) {
	room.Connections = append(room.Connections, connectingRoom)
	connectingRoom.Connections = append(connectingRoom.Connections, room)
}