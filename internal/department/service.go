package department

import (
	"errors"

	"github.com/google/uuid"
)

type DepartmentService struct {
	repo *DepartmentRepository
}

func NewDepartmentService(repo *DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (srv *DepartmentService) Create(params *DepartmentCreateRequest) (*Department, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	data := &Department{
		ID:          uuid.String(),
		Name:        params.Name,
		Description: params.Description}
	exist, err := srv.repo.ExistsByName(params.Name)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("部门已存在")
	}
	err = srv.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *DepartmentService) Detail(id string) (*Department, error) {
	data, err := srv.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
