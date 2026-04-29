package user

import (
	"context"
	"errors"
	"fmt"
	"go-learn/internal/department"
	"go-learn/internal/role"

	"github.com/google/uuid"
)

type UserService struct {
	repo      *UserRepository
	depUtils  *department.Utils
	roleUtils *role.Utils
}

func NewUserService(repo *UserRepository, depUtils *department.Utils, roleUtils *role.Utils) *UserService {
	return &UserService{repo: repo, depUtils: depUtils, roleUtils: roleUtils}
}
func (s *UserService) Register(ctx context.Context, p *UserRegisterReq) (*User, error) {
	exist, err := s.repo.ExistsByEmail(ctx, p.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("邮箱已注册")
	}

	done := make(chan error, 2) // 缓冲 2，保证两个 goroutine 都能写入不阻塞
	var dep *department.Department
	var roleResult *role.Role

	go func() {
		var err error
		dep, err = s.depUtils.CheckID(ctx, p.DepartmentID)
		done <- err
	}()
	go func() {
		var err error
		roleResult, err = s.roleUtils.CheckID(ctx, p.RoleID)
		done <- err
	}()

	for range 2 {
		if err := <-done; err != nil {
			return nil, fmt.Errorf("校验失败: %w", err)
		}
	}

	if dep == nil || roleResult == nil {
		return nil, errors.New("部门或角色不存在")
	}
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:           uuid.String(),
		Name:         p.Name,
		Email:        p.Email,
		DepartmentID: p.DepartmentID,
		RoleID:       p.RoleID,
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Update(ctx context.Context, id string, p *UserUpdateReq) (*User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if p.Email != user.Email {
		exist, err := s.repo.ExistsByEmail(ctx, p.Email)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, errors.New("邮箱已注册")
		}
	}
	user.Email = p.Email
	user.Name = p.Name
	err = s.repo.Update(ctx, id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
