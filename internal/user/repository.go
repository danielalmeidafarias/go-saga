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
	Nickname *string `json:"nickname,omitempty"`
	Email    *string `json:"email,omitempty"`
}

func (r *UserRepository) Create(in CreateUserInput) (*User, error) {
	user := User{
		Name:     in.Name,
		Nickname: in.Nickname,
		Email:    in.Email,
		Birthday: in.Birthday,
		Passport: in.Passport,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetOneByUUID(uuid uuid.UUID) (*User, error) {
	var user User

	err := r.db.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(uuid uuid.UUID, in UpdateUserInput) (*User, error) {
	user, err := r.GetOneByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(user).Updates(in).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Delete(uuid uuid.UUID) error {
	user, err := r.GetOneByUUID(uuid)
	if err != nil {
		return err
	}

	if err := r.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
