package room

type roomService struct {
	roomRepository IRoomRepository
	uRRepository   RoomUsersRepository
}

type IRoomService interface {
	GetRoom(id string) (map[string]interface{}, error)
	GetAllRooms(userID string, limit, offset int64) ([]map[string]interface{}, error)
	CreateRoom(room map[string]interface{}) (map[string]interface{}, error)
	DeleteRoom(roomID string) error
	UpdateRoom(room map[string]interface{}) (map[string]interface{}, error)
	AddUsers(roomID string, users []string) (err error)
	DeleteUsers(roomID string, users []string) (err error)
}

func (s *roomService) GetRoom(id string) (map[string]interface{}, error) {
	result, err := s.roomRepository.GetRoom(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *roomService) GetAllRooms(userID string, limit, offset int64) ([]map[string]interface{}, error) {
	rooms, err := s.uRRepository.GetRoomsByUserId(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range rooms {
		rooms[i]["id"] = rooms[i]["_id"]
		delete(rooms[i], "_id")
	}

	return rooms, nil
}

func (s *roomService) CreateRoom(room map[string]interface{}) (map[string]interface{}, error) {
	result, err := s.roomRepository.CreateRoom(room)
	if err != nil {
		return nil, err
	}
	_, err = s.uRRepository.CreateUserRoomConnection(map[string]interface{}{
		"user_id": result["created_by"],
		"room_id": result["_id"],
	})
	if err != nil {
		return nil, err
	}

	result["id"] = result["_id"]
	delete(result, "_id")
	return result, nil
}

func (s *roomService) DeleteRoom(roomID string) error {
	return s.roomRepository.DeleteRoom(roomID)
}

func (s *roomService) UpdateRoom(room map[string]interface{}) (map[string]interface{}, error) {
	updateResult, err := s.roomRepository.UpdateRoom(room)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (s *roomService) AddUsers(roomId string, users []string) (err error) {
	for _, user := range users {
		_, err = s.uRRepository.CreateUserRoomConnection(map[string]interface{}{
			"user_id": user,
			"room_id": roomId,
		})
		if err != nil {
			return
		}
	}
	return
}

func (s *roomService) DeleteUsers(roomId string, users []string) (err error) {
	for _, user := range users {
		err = s.uRRepository.DeleteUserRoomConnection(map[string]interface{}{
			"user_id": user,
			"room_id": roomId,
		})
		if err != nil {
			return
		}
	}
	return
}

func NewRoomService(roomRepository IRoomRepository, uRRepository RoomUsersRepository) IRoomService {
	return &roomService{
		roomRepository,
		uRRepository,
	}
}
