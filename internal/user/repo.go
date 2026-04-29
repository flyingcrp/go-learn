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
