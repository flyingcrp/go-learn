package department

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type DepartmentRepo interface {
	FindByID(ctx context.Context, id string) (*Department, error)
	Create(ctx context.Context, department *Department) error
	ExistsByName(ctx context.Context, name string) (bool, error)
}
type DepartmentService struct {
	repo DepartmentRepo
}

func NewDepartmentService(repo DepartmentRepo) *DepartmentService {
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
func (srv *DepartmentService) CheckID(ctx context.Context, id string) (*Department, error) {
	return srv.repo.FindByID(ctx, id)
}
