package entity

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
