package user

import "errors"

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
		return nil, errors.New("邮箱已被注册")
	}
	user := &User{
		ID:           "generateID()1",
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
