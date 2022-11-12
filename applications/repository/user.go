package repository

import (
	"github.com/destafajri/auth-app/applications/entity"
)

type UserInterface interface {
	Register(*entity.UserEntity) error
	Remindme(string)(*entity.UserEntity, error)
	Login(string) (*entity.UserEntity, error)
}