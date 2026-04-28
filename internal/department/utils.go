package department

type Utils struct {
	repo *DepartmentRepository
}

func NewUtils(repo *DepartmentRepository) *Utils {
	return &Utils{repo: repo}
}

func (u *Utils) CheckID(id string) (*Department, error) {
	return u.repo.FindByID(id)
}
