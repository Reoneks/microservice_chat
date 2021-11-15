package service

import (
	"chatex/room/room/repository"
	room "chatex/room/room/room_interface"
	userRep "chatex/room/room_user/repository"
	room_user "chatex/room/room_user/room_user_interface"
	"errors"
)

type RoomServiceImpl struct {
	roomRepository room.RoomRepository
	uRRepository   userRep.RoomUsersRepository
}

func (s *RoomServiceImpl) GetRoom(id string) (*room.Rooms, error) {
	result, err := s.roomRepository.GetRoom(id)
	if err != nil {
		return nil, err
	}
	resultRoom := repository.FromRoomsDto(*result)
	return &resultRoom, nil
}

func (s *RoomServiceImpl) GetRooms(filter *room.RoomsFilter) ([]room.Rooms, error) {
	result, err := s.roomRepository.GetRooms(filter)
	if err != nil {
		return nil, err
	}
	resultRooms := repository.FromRoomsDtos(result)
	return resultRooms, nil
}

func (s *RoomServiceImpl) CreateRoom(room *room.Rooms) (*room.Rooms, error) {
	result, err := s.roomRepository.CreateRoom(repository.ToRoomsDto(*room))
	if err != nil {
		return nil, err
	}
	_, err = s.uRRepository.CreateUserRoomConnection(room_user.RoomUsersDto{
		UserID: result.CreatedBy,
		RoomID: result.ID,
	})
	if err != nil {
		return nil, err
	}
	resultMessages := repository.FromRoomsDto(*result)
	return &resultMessages, nil
}

func (s *RoomServiceImpl) DeleteRoom(Room_id, userId string) error {
	result, err := s.GetRoom(Room_id)
	if err != nil {
		return err
	} else if result.CreatedBy != userId {
		return errors.New("you are not allowed to do it")
	}
	return s.roomRepository.DeleteRoom(Room_id)
}

func (s *RoomServiceImpl) UpdateRoom(room *room.Rooms, userId string) (*room.Rooms, error) {
	result, err := s.GetRoom(room.ID)
	if err != nil {
		return nil, err
	} else if result.CreatedBy != userId {
		return nil, errors.New("you are not allowed to do it")
	}
	updateResult, err := s.roomRepository.UpdateRoom(repository.ToRoomsDto(*room))
	if err != nil {
		return nil, err
	}
	resultRoom := repository.FromRoomsDto(*updateResult)
	return &resultRoom, nil
}

func (s *RoomServiceImpl) AddUsers(roomId, userId string, users []string) (err error) {
	result, err := s.GetRoom(roomId)
	if err != nil {
		return
	} else if result.CreatedBy != userId {
		err = errors.New("you are not allowed to do it")
		return
	}
	for _, user := range users {
		_, err = s.uRRepository.CreateUserRoomConnection(room_user.RoomUsersDto{
			UserID: user,
			RoomID: roomId,
		})
		if err != nil {
			return
		}
	}
	return
}

func NewRoomService(roomRepository room.RoomRepository, uRRepository userRep.RoomUsersRepository) room.RoomService {
	return &RoomServiceImpl{
		roomRepository,
		uRRepository,
	}
}
