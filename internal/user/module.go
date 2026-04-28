package user

import (
	"go-learn/internal/common/storage"
	"go-learn/internal/department"
)

func NewUserModule(infra *storage.Infra, depUtils *department.Utils) *UserHandler {
	repo := NewUserRepository(infra.MySQL)
	srv := NewUserService(repo, depUtils)
	return &UserHandler{srv: srv}
}
