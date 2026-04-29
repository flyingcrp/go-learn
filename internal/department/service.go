package department

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type DepartmentService struct {
	repo *DepartmentRepository
}

func NewDepartmentService(repo *DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (srv *DepartmentService) Create(ctx context.Context, params *DepartmentCreateRequest) (*Department, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	data := &Department{
		ID:          uuid.String(),
		Name:        params.Name,
		Description: params.Description}
	exist, err := srv.repo.ExistsByName(ctx, params.Name)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("部门已存在")
	}
	err = srv.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *DepartmentService) Detail(ctx context.Context, id string) (*Department, error) {
	data, err := srv.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
