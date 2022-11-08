package service

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type IUserService interface {
	Create(ctx context.Context, r *User) (*User, error)
	Update(ctx context.Context, r *User) (*User, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
}

type userService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) IUserService {
	return &userService{
		db,
	}
}

func (s *userService) Create(ctx context.Context, r *User) (c *User, err error) {
	return nil, nil
}
func (s *userService) Update(ctx context.Context, r *User) (c *User, err error) {
	return nil, nil
}
func (s *userService) Delete(ctx context.Context, id int) (err error) {
	return nil
}
func (s *userService) FindAll(ctx context.Context) (c []*User, err error) {
	return nil, nil
}
func (s *userService) FindByID(ctx context.Context, id int) (c *User, err error) {
	return nil, nil
}
