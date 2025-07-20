package user

import (
	"context"
	"time"

	"github.com/danielalmeidafarias/go-saga/pkg"
	"github.com/google/uuid"
)

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

	input := CreateUserInput{
		Name:     params.Name,
		Nickname: params.Nickname,
		Email:    params.Email,
		Birthday: birthday,
		Passport: params.Passport,
	}

	user, err := s.userRepository.Create(input)
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{
		UserUUID: user.UUID.String(),
		Message:  "User created successfully",
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*GetUserResponse, error) {
	userUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.GetOneByUUID(userUUID)
	if err != nil {
		return nil, err
	}

	return &GetUserResponse{
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

	userUUID, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.Update(userUUID, UpdateUserInput{
		Nickname: params.Nickname,
		Email:    params.Email,
	})
	if err != nil {
		return nil, err
	}

	return &UpdateUserResponse{
		Message: "User updated successfully",
		User: &GetUserResponse{
			UUID:     user.UUID.String(),
			Name:     user.Name,
			Nickname: user.Nickname,
			Email:    user.Email,
			Birthday: user.Birthday,
			Passport: user.Passport,
		},
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) (*DeleteUserResponse, error) {
	userUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if err := s.userRepository.Delete(userUUID); err != nil {
		return nil, err
	}

	return &DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
