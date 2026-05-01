package department

import "context"

type Utils struct {
	repo *gormDepartmentRepo
}

func NewUtils(repo *gormDepartmentRepo) *Utils {
	return &Utils{repo: repo}
}

func (u *Utils) CheckID(ctx context.Context, id string) (*Department, error) {
	return u.repo.FindByID(ctx, id)
}
