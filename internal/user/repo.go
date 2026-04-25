package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var cnt int64
	err := r.db.Model(&User{}).Where("email = ?", email).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}
func (r *UserRepository) Create(user *User) error {
	return r.db.Create(user).Error
}
func (r *UserRepository) Update(id string, newUser *User) error {
	return r.db.Where("id = ?", id).Updates(newUser).Error
}
func (r *UserRepository) FindByID(id string) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).Omit("Token").Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
