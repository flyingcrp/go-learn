package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
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

	user := &User{
		ID:           uuid.String(),
		Name:         p.Name,
		Email:        p.Email,
		DepartmentID: "p.DepartmentID",
		RoleID:       "p.RoleID",
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
