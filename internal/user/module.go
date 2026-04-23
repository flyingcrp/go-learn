package user

import "go-learn/internal/common/storage"

func NewUserModule(infra *storage.Infra) *UserHandler {
	repo := NewUserRepository(infra.MySQL)
	srv := NewUserService(repo)
	return &UserHandler{srv: srv}
}
