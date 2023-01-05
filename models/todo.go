package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"author" example:"sport"`
	Description string `json:"description" example:"play football"`
	HasDone     string `json:"has_done" example:"done"`
}
