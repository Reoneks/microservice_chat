package room

import (
	"time"
)

type Rooms struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    int64     `json:"status"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
