package role

import "gorm.io/gorm"

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (repo *RoleRepository) FindByID(id string) (*Role, error) {
	var role Role
	err := repo.db.Where("id = ?", id).First(&role).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (repo *RoleRepository) ExistsByName(name string) (bool, error) {
	var cnt int64
	err := repo.db.Model(&Role{}).Where("name = ?", name).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (repo *RoleRepository) ExistsByCode(code string) (bool, error) {
	var cnt int64
	err := repo.db.Model(&Role{}).Where("code = ?", code).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (repo *RoleRepository) Create(role *Role) error {
	return repo.db.Create(role).Error
}
