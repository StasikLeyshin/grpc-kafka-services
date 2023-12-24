package models

type User struct {
	UUID string `json:"uuid" gorm:"primaryKey"`
}
