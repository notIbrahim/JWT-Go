package entity

import "time"

type Model struct {
	ID        uint       `gorm:"primary" json:"ID"`
	CreatedAt *time.Time `json:"created_at, omitempty"`
	UpdatedAt *time.Time `json:"updated_at, omitempty"`
}
