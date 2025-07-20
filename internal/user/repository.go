package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type CreateUserInput struct {
	Name     string
	Nickname *string
	Email    string
	Birthday time.Time
	Passport string
}

type UpdateUserInput struct {
	Nickname *string
	Email    *string
}

func (r *UserRepository) Create(input CreateUserInput) (*User, error) {
	user := User{
		Name:     input.Name,
		Nickname: input.Nickname,
		Email:    input.Email,
		Birthday: input.Birthday,
		Passport: input.Passport,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetOneById(id string) (*User, error) {
	var user User

	if userUUID, err := uuid.Parse(id); err == nil {
		if err := r.db.Where("uuid = ?", userUUID).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("user not found")
			}
			return nil, err
		}
	} else {
		if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("user not found")
			}
			return nil, err
		}
	}

	return &user, nil
}

func (r *UserRepository) Update(id string, input UpdateUserInput) (*User, error) {
	user, err := r.GetOneById(id)
	if err != nil {
		return nil, err
	}

	updateData := make(map[string]interface{})

	if input.Nickname != nil {
		updateData["nickname"] = *input.Nickname
	}

	if input.Email != nil {
		updateData["email"] = *input.Email
	}

	if len(updateData) > 0 {
		if err := r.db.Model(user).Updates(updateData).Error; err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *UserRepository) Delete(id string) error {
	user, err := r.GetOneById(id)
	if err != nil {
		return err
	}

	if err := r.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
