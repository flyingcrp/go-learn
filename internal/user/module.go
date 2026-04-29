package user

import (
	"go-learn/internal/common/storage"
	"go-learn/internal/department"
	"go-learn/internal/role"
)

func NewUserModule(infra *storage.Infra, depUtils *department.Utils, roleUtils *role.Utils) *UserHandler {
	repo := NewUserRepository(infra.MySQL)
	srv := NewUserService(repo, depUtils, roleUtils)
	return &UserHandler{srv: srv}
}
