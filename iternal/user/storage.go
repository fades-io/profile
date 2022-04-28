package user

import "context"

type Repository interface {
	Create(ctx context.Context, author *User) error
	// do we need it?
	FindAll(ctx context.Context) (u []User, err error)
	FindOne(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
