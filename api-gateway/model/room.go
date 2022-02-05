package model

type Pagination struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"next_offset"`
}

type PaginationRoomsResponse struct {
	Pagination
	Rooms []map[string]interface{} `json:"rooms"`
}

type UserIDsRequest struct {
	UserIDs []string `json:"user_ids"`
}
