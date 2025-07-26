package domain

import "context"

type UserRepository interface {
	GetUser(ctx context.Context, userID string) *Error
}
