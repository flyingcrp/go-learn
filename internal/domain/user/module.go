package user

import (
	"go-learn/internal/infra/event"
	"go-learn/internal/infra/storage"
)

func NewUserModule(infra *storage.Infra, depUtils departmentChecker, roleChecker roleChecker, bus *event.Bus, jwtSecret string) *UserHandler {
	repo := NewUserRepository(infra.MySQL)
	srv := NewUserService(repo, depUtils, roleChecker, bus, jwtSecret)
	return &UserHandler{srv: srv}
}
