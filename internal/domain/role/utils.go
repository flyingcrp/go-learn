package role

import "context"

type Utils struct {
	repo *gormRoleRepo
}

func NewUtils(repo *gormRoleRepo) *Utils {
	return &Utils{repo: repo}
}

func (u *Utils) CheckID(ctx context.Context, id string) (*Role, error) {
	return u.repo.FindByID(ctx, id)
}
