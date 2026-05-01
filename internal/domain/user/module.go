package user

import (
	"go-learn/internal/domain/department"
	"go-learn/internal/domain/role"
	"go-learn/internal/infra/storage"
)

func NewUserModule(infra *storage.Infra, depUtils *department.Utils, roleUtils *role.Utils, jwtSecret string) *UserHandler {
	repo := NewUserRepository(infra.MySQL)
	srv := NewUserService(repo, depUtils, roleUtils, jwtSecret)
	return &UserHandler{srv: srv}
}
