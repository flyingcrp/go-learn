package department

import "context"

type Utils struct {
	repo *DepartmentRepository
}

func NewUtils(repo *DepartmentRepository) *Utils {
	return &Utils{repo: repo}
}

func (u *Utils) CheckID(ctx context.Context, id string) (*Department, error) {
	return u.repo.FindByID(ctx, id)
}
