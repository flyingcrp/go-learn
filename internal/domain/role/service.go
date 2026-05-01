package role

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type RoleRepo interface {
	ExistsByName(ctx context.Context, name string) (bool, error)
	ExistsByCode(ctx context.Context, code string) (bool, error)
	Create(ctx context.Context, role *Role) error
	FindByID(ctx context.Context, id string) (*Role, error)
}

type RoleService struct {
	repo RoleRepo
}

func NewRoleService(repo RoleRepo) *RoleService {
	return &RoleService{repo: repo}
}

func (srv *RoleService) Create(ctx context.Context, params *RoleCreateRequest) (*Role, error) {
	exist, err := srv.repo.ExistsByName(ctx, params.Name)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("角色名称已存在")
	}

	exist, err = srv.repo.ExistsByCode(ctx, params.Code)
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
	if err := srv.repo.Create(ctx, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (srv *RoleService) Detail(ctx context.Context, id string) (*Role, error) {
	return srv.repo.FindByID(ctx, id)
}
