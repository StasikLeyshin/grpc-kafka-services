package model

import "time"

type Server struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	Port      string    `json:"port"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateServerRequest struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type CreateServerResponse struct {
	Status string `json:"ok"`
}
