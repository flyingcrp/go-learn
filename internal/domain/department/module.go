package department

import "go-learn/internal/infra/storage"

type DepartmentModule struct {
	Handler *DepartmentHandler
	Utils   *Utils
}

func NewDepartmentModule(infra *storage.Infra) *DepartmentModule {
	repo := NewDepartmentRepository(infra.MySQL)
	srv := NewDepartmentService(repo)
	utils := NewUtils(repo)
	return &DepartmentModule{Handler: &DepartmentHandler{srv: srv}, Utils: utils}
}
