package department

import "gorm.io/gorm"

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (repo *DepartmentRepository) FindByID(id string) (*Department, error) {
	var dep Department
	err := repo.db.Where("id = ?", id).First(&dep).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &dep, nil
}
func (repo *DepartmentRepository) ExistsByName(name string) (bool, error) {
	var cnt int64
	err := repo.db.Model(&Department{}).Where("name = ?", name).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}
func (repo *DepartmentRepository) Create(department *Department) error {
	return repo.db.Create(department).Error
}
