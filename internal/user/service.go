package user

import (
	"errors"
	"fmt"
	"go-learn/internal/department"

	"github.com/google/uuid"
)

type UserService struct {
	repo     *UserRepository
	depUtils *department.Utils
}

func NewUserService(repo *UserRepository, depUtils *department.Utils) *UserService {
	return &UserService{repo: repo, depUtils: depUtils}
}
func (s *UserService) Register(p *UserRegisterReq) (*User, error) {
	exist, err := s.repo.ExistsByEmail(p.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("邮箱已注册")
	}
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	dep, err := s.depUtils.CheckID(p.DepartmentID)
	if err != nil {
		return nil, err
	}
	if dep == nil {
		return nil, errors.New("无效的部门 ID")
	}

	user := &User{
		ID:           uuid.String(),
		Name:         p.Name,
		Email:        p.Email,
		DepartmentID: p.DepartmentID,
		RoleID:       p.RoleID,
	}
	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Update(id string, p *UserUpdateReq) (*User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if p.Email != user.Email {
		exist, err := s.repo.ExistsByEmail(p.Email)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, errors.New("邮箱已注册")
		}
	}
	user.Email = p.Email
	user.Name = p.Name
	err = s.repo.Update(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
