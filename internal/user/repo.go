package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var cnt int64
	err := r.db.WithContext(ctx).Model(&User{}).Where("email = ?", email).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}
func (r *UserRepository) Create(ctx context.Context, user *User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
func (r *UserRepository) Update(ctx context.Context, id string, newUser *User) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Updates(newUser).Error
}
func (r *UserRepository) FindByID(ctx context.Context, id string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("id = ?", id).Omit("Token").Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Login(ctx context.Context, name, email string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("name = ? AND email = ?", name, email).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) List(ctx context.Context, params *UserListReq) (list []User, total int64, err error) {
	baseQuery := r.db.WithContext(ctx).Model(&User{})
	if params.Name != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if err = baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return []User{}, 0, nil
	}
	if err = baseQuery.
		Limit(params.PageSize).
		Offset((params.Page - 1) * params.PageSize).
		Omit("Token").
		Order("created_at desc").
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
