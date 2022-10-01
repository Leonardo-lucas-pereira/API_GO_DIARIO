package models

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsAdm      bool      `json:"is_adm"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"updated"`
	DeleteddAt time.Time `json:"deleted" gorm:"index"`
}
