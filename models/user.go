package models

import "time"

type UserDTO struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username"  validate:"required,min=3,max=32"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
	Age      int    `json:"age" validate:"omitempty,gte=13,lte=99"`
	Role     string `json:"role" validate:"required,oneof=user creator"`
}

func (u *UserDTO) ToEntity() User {
	return User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Age:      u.Age,
		Role:     u.Role,
	}
}

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Username  string    `bson:"username"`
	Email     string    `bson:"email" `
	Password  string    `bson:"password"`
	Age       int       `bson:"age"`
	Role      string    `bson:"role"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
		Age:      u.Age,
		Role:     u.Role,
	}
}
