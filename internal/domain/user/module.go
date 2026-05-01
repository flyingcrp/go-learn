package user

import (
	"go-learn/internal/infra/storage"
)

func NewUserModule(infra *storage.Infra, depUtils DepartmentChecker, roleChecker RoleChecker, jwtSecret string) *UserHandler {
	repo := NewUserRepository(infra.MySQL)
	srv := NewUserService(repo, depUtils, roleChecker, jwtSecret)
	return &UserHandler{srv: srv}
}
