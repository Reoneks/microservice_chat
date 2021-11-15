package room

type RoomRepository interface {
	GetRoom(id string) (*RoomsDto, error)
	CreateRoom(user RoomsDto) (*RoomsDto, error)
	UpdateRoom(user RoomsDto) (*RoomsDto, error)
	DeleteRoom(id string) error
	GetRooms(filter *RoomsFilter) ([]RoomsDto, error)
}

type RoomService interface {
	GetRoom(id string) (*Rooms, error)
	GetRooms(filter *RoomsFilter) ([]Rooms, error)
	CreateRoom(room *Rooms) (*Rooms, error)
	DeleteRoom(room_id, userId string) error
	UpdateRoom(room *Rooms, userId string) (*Rooms, error)
	AddUsers(roomId, userId string, users []string) (errorsArray error)
}
