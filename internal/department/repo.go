package department

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (repo *DepartmentRepository) FindByID(ctx context.Context, id string) (*Department, error) {
	var dep Department
	err := repo.db.WithContext(ctx).Where("id = ?", id).First(&dep).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &dep, nil
}
func (repo *DepartmentRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var cnt int64
	err := repo.db.WithContext(ctx).Model(&Department{}).Where("name = ?", name).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}
func (repo *DepartmentRepository) Create(ctx context.Context, department *Department) error {
	return repo.db.WithContext(ctx).Create(department).Error
}
