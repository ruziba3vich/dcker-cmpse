package models

import "time"

type (
	GetRequest struct {
		Message    string    `json:"message"`
		ReceivedAt time.Time `json:"received_at"`
	}

	Response struct {
		Message    string    `json:"message"`
		ReceivedAt time.Time `json:"received_at"`
		ResentAt   time.Time `json:"resent_at"`
	}
)

func (g *GetRequest) Convert() *Response {
	var response Response
	response.Message = g.Message
	response.ReceivedAt = g.ReceivedAt
	return &response
}
