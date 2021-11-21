package room

import (
	"errors"

	"github.com/Reoneks/microservice_chat/room/model"
	room_user "github.com/Reoneks/microservice_chat/room/room_user"
)

type roomService struct {
	roomRepository IRoomRepository
	uRRepository   room_user.RoomUsersRepository
}

type IRoomService interface {
	GetRoom(id string) (*model.RoomsDto, error)
	GetRooms(filter *RoomsFilter) ([]model.RoomsDto, error)
	CreateRoom(room *model.RoomsDto) (*model.RoomsDto, error)
	DeleteRoom(room_id, userId string) error
	UpdateRoom(room *model.RoomsDto, userId string) (*model.RoomsDto, error)
	AddUsers(roomId, userId string, users []string) (errorsArray error)
}

func (s *roomService) GetRoom(id string) (*model.RoomsDto, error) {
	result, err := s.roomRepository.GetRoom(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *roomService) GetRooms(filter *RoomsFilter) ([]model.RoomsDto, error) {
	result, err := s.roomRepository.GetRooms(filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *roomService) CreateRoom(room *model.RoomsDto) (*model.RoomsDto, error) {
	result, err := s.roomRepository.CreateRoom(room)
	if err != nil {
		return nil, err
	}
	_, err = s.uRRepository.CreateUserRoomConnection(model.RoomUsersDto{
		UserID: result.CreatedBy,
		RoomID: result.ID,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *roomService) DeleteRoom(room_id, userId string) error {
	result, err := s.GetRoom(room_id)
	if err != nil {
		return err
	} else if result.CreatedBy != userId {
		return errors.New("you are not allowed to do it")
	}
	return s.roomRepository.DeleteRoom(room_id)
}

func (s *roomService) UpdateRoom(room *model.RoomsDto, userId string) (*model.RoomsDto, error) {
	result, err := s.GetRoom(room.ID)
	if err != nil {
		return nil, err
	} else if result.CreatedBy != userId {
		return nil, errors.New("you are not allowed to do it")
	}

	updateResult, err := s.roomRepository.UpdateRoom(room)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (s *roomService) AddUsers(roomId, userId string, users []string) (err error) {
	result, err := s.GetRoom(roomId)
	if err != nil {
		return
	} else if result.CreatedBy != userId {
		err = errors.New("you are not allowed to do it")
		return
	}
	for _, user := range users {
		_, err = s.uRRepository.CreateUserRoomConnection(model.RoomUsersDto{
			UserID: user,
			RoomID: roomId,
		})
		if err != nil {
			return
		}
	}
	return
}

func NewRoomService(roomRepository IRoomRepository, uRRepository room_user.RoomUsersRepository) IRoomService {
	return &roomService{
		roomRepository,
		uRRepository,
	}
}
