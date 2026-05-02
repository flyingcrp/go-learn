package user

import (
	"context"
	"errors"
	"fmt"
	"go-learn/internal/domain/department"
	"go-learn/internal/domain/role"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type UserRepo interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id string, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	Login(ctx context.Context, name, email string) (*User, error)
	List(ctx context.Context, params *UserListReq) ([]User, int64, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type DepartmentChecker interface {
	CheckID(ctx context.Context, id string) (*department.Department, error)
}
type RoleChecker interface {
	CheckID(ctx context.Context, id string) (*role.Role, error)
}

type UserService struct {
	repo        UserRepo
	depChecker  DepartmentChecker
	roleChecker RoleChecker
	jwtSecret   string
}

func NewUserService(repo UserRepo, depUtils DepartmentChecker, roleUtils RoleChecker, jwtSecret string) *UserService {
	return &UserService{repo: repo, depChecker: depUtils, roleChecker: roleUtils, jwtSecret: jwtSecret}
}
func (s *UserService) Register(ctx context.Context, p *UserRegisterReq) (*User, error) {
	exist, err := s.repo.ExistsByEmail(ctx, p.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("邮箱已注册")
	}
	g, egCtx := errgroup.WithContext(ctx)
	egCtx, cancel := context.WithTimeout(egCtx, 3*time.Second)
	defer cancel()
	var dep *department.Department
	var roleResult *role.Role

	g.Go(func() error {
		d, err := s.depChecker.CheckID(egCtx, p.DepartmentID)
		if err != nil {
			return err
		}
		dep = d
		return nil
	})
	g.Go(func() error {
		r, err := s.roleChecker.CheckID(egCtx, p.RoleID)
		if err != nil {
			return err
		}
		roleResult = r
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("校验失败: %w", err)
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

func (s *UserService) FindByID(ctx context.Context, id string) (*User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
func (s *UserService) Login(ctx context.Context, param *UserLoginReq) (*User, error) {
	user, err := s.repo.Login(ctx, param.Name, param.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	return user, nil
}
func (s *UserService) GenerateJWT(user *User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(s.jwtSecret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *UserService) List(ctx context.Context, params *UserListReq) (*UserListResp, error) {
	list, total, e := s.repo.List(ctx, params)
	if e != nil {
		return nil, e
	}
	items := make([]UserDetailResp, len(list))
	for i, u := range list {
		items[i] = *toUserDetailResp(&u)
	}
	return &UserListResp{
		List:  items,
		Total: total,
	}, nil
}
