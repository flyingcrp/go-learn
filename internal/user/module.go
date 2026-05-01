package user

import (
	"go-learn/internal/department"
	"go-learn/internal/role"
	"go-learn/internal/storage"
)

func NewUserModule(infra *storage.Infra, depUtils *department.Utils, roleUtils *role.Utils, jwtSecret string) *UserHandler {
	repo := NewUserRepository(infra.MySQL)
	srv := NewUserService(repo, depUtils, roleUtils, jwtSecret)
	return &UserHandler{srv: srv}
}
