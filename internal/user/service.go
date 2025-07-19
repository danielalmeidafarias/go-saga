package user

import (
	"context"
	"fmt"
	"time"

	"github.com/danielalmeidafarias/go-saga/pkg"
)

type CreateUserParams struct {
	Name     string `json:"name" validate:"required"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Birthday string `json:"birthday" validate:"required"`
	Passport string `json:"passport" validate:"required,len=8"`
}

type CreateUserResponse struct {
	UserID   string `json:"userId"`
	UserUUID string `json:"userUuid"`
	Message  string `json:"message"`
}

type GetUserResponse struct {
	ID       uint      `json:"id"`
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Nickname *string   `json:"nickname,omitempty"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Passport string    `json:"passport"`
}

type UpdateUserParams struct {
	ID       string `json:"id" validate:"required"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
}

type UpdateUserResponse struct {
	Message string           `json:"message"`
	User    *GetUserResponse `json:"user"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

type UserService struct {
	validator      *pkg.Validator
	userRepository *UserRepository
}

func NewUserService(
	userRepository *UserRepository,
	validator *pkg.Validator,
) *UserService {
	return &UserService{
		validator:      validator,
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, params CreateUserParams) (*CreateUserResponse, error) {
	if err := s.validator.Validate(params); err != nil {
		return nil, err
	}

	birthday, err := time.Parse("2006-01-02", params.Birthday)
	if err != nil {
		return nil, err
	}

	var nickname *string
	if params.Nickname != "" {
		nickname = &params.Nickname
	}

	input := CreateUserInput{
		Name:     params.Name,
		Nickname: nickname,
		Email:    params.Email,
		Birthday: birthday,
		Passport: params.Passport,
	}

	user, err := s.userRepository.Create(input)
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{
		UserID:   fmt.Sprintf("%d", user.ID),
		UserUUID: user.UUID.String(),
		Message:  "User created successfully",
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*GetUserResponse, error) {
	user, err := s.userRepository.GetOneById(id)
	if err != nil {
		return nil, err
	}

	return &GetUserResponse{
		ID:       user.ID,
		UUID:     user.UUID.String(),
		Name:     user.Name,
		Nickname: user.Nickname,
		Email:    user.Email,
		Birthday: user.Birthday,
		Passport: user.Passport,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, params UpdateUserParams) (*UpdateUserResponse, error) {
	if err := s.validator.Validate(params); err != nil {
		return nil, err
	}

	var nickname *string
	if params.Nickname != "" {
		nickname = &params.Nickname
	}

	var email *string
	if params.Email != "" {
		email = &params.Email
	}

	input := UpdateUserInput{
		Nickname: nickname,
		Email:    email,
	}

	user, err := s.userRepository.Update(params.ID, input)
	if err != nil {
		return nil, err
	}

	userResponse := &GetUserResponse{
		ID:       user.ID,
		UUID:     user.UUID.String(),
		Name:     user.Name,
		Nickname: user.Nickname,
		Email:    user.Email,
		Birthday: user.Birthday,
		Passport: user.Passport,
	}

	return &UpdateUserResponse{
		Message: "User updated successfully",
		User:    userResponse,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) (*DeleteUserResponse, error) {
	if err := s.userRepository.Delete(id); err != nil {
		return nil, err
	}

	return &DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
