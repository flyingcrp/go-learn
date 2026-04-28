package role

import (
	"errors"

	"github.com/google/uuid"
)

type RoleService struct {
	repo *RoleRepository
}

func NewRoleService(repo *RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (srv *RoleService) Create(params *RoleCreateRequest) (*Role, error) {
	exist, err := srv.repo.ExistsByName(params.Name)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("角色名称已存在")
	}

	exist, err = srv.repo.ExistsByCode(params.Code)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("角色编码已存在")
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	role := &Role{
		ID:          id.String(),
		Name:        params.Name,
		Code:        params.Code,
		Description: params.Description,
	}
	if err := srv.repo.Create(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (srv *RoleService) Detail(id string) (*Role, error) {
	return srv.repo.FindByID(id)
}
