package model

type PaginationMessagesResponse struct {
	Pagination
	Messages []map[string]interface{} `json:"messages"`
}
