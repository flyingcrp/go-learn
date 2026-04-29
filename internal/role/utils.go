package role

type Utils struct {
	repo *RoleRepository
}

func NewUtils(repo *RoleRepository) *Utils {
	return &Utils{repo: repo}
}

func (u *Utils) CheckID(id string) (*Role, error) {
	return u.repo.FindByID(id)
}
