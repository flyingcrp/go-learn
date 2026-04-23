package user

import (
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
