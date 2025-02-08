package models

import (
	"pioApi/ent/schema"
)

type User struct {
	schema.User
	// elements used for post requests
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
