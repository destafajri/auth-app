package repository

import (
	"github.com/destafajri/auth-app/applications/entity"
)

type UserInterface interface {
	Register(*entity.UserEntity) error
}