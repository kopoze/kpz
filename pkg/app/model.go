package app

import "time"

type App struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Subdomain string    `json:"subdomain"`
	Port      string    `json:"port"`
}
