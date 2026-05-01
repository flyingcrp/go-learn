package role

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type gormRoleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *gormRoleRepo {
	return &gormRoleRepo{db: db}
}

func (repo *gormRoleRepo) FindByID(ctx context.Context, id string) (*Role, error) {
	var role Role
	err := repo.db.WithContext(ctx).Where("id = ?", id).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (repo *gormRoleRepo) ExistsByName(ctx context.Context, name string) (bool, error) {
	var cnt int64
	err := repo.db.WithContext(ctx).Model(&Role{}).Where("name = ?", name).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (repo *gormRoleRepo) ExistsByCode(ctx context.Context, code string) (bool, error) {
	var cnt int64
	err := repo.db.WithContext(ctx).Model(&Role{}).Where("code = ?", code).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (repo *gormRoleRepo) Create(ctx context.Context, role *Role) error {
	return repo.db.WithContext(ctx).Create(role).Error
}
