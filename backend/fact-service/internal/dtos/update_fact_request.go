package dtos

type UpdateFactRequest struct {
	ID   uint64    `json:"id"`
	Text string `json:"text"`
}
