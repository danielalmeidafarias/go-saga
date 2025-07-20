package user

import "time"

type CreateUserParams struct {
	Name     string  `json:"name" validate:"required"`
	Nickname *string `json:"nickname,omitempty"`
	Email    string  `json:"email" validate:"required,email"`
	Birthday string  `json:"birthday" validate:"required"`
	Passport string  `json:"passport" validate:"required,len=8"`
}

type UpdateUserParams struct {
	ID       string  `json:"id" validate:"required"`
	Nickname *string `json:"nickname,omitempty"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
}

type CreateUserResponse struct {
	UserUUID string `json:"userUuid"`
	Message  string `json:"message"`
}

type GetUserResponse struct {
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Nickname *string   `json:"nickname,omitempty"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Passport string    `json:"passport"`
}

type UpdateUserResponse struct {
	Message string           `json:"message"`
	User    *GetUserResponse `json:"user"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
