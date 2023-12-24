package models

import "time"

type Server struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	Port      string    `json:"port"`
	CreatedAt time.Time `json:"created_at"`
}
