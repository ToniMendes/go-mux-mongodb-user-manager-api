package repository

import (
	"go-mux-mongodb-user-manager-api/internal/usecases/users_manager"
)

type WriterOnlyRepositoryUseCases interface {
	ExecCreate(users_manager.UserCreateInput) (users_manager.UserCreateResponse, error)
	ExecUpdateName(users_manager.UserUpdateNameInput) (map[string]interface{}, error)
	ExecUpdateEmail(users_manager.UserUpdateEmailInput) (map[string]interface{}, error)
}

type ReadOnlyRepositoryUseCases interface {
	ExecGetAll() ([]users_manager.UserGetAllResponse, error)
	ExecLogin(users_manager.UserLoginInput) (users_manager.UserGetByEmailResponse, error)
}
