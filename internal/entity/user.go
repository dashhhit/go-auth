package entity

import "time"

type User struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	Birthday  time.Time `json:"birthday"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Telephone string    `json:"telephone"`
}
