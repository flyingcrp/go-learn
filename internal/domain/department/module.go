package department

import "go-learn/internal/infra/storage"

type DepartmentModule struct {
	Handler *DepartmentHandler
	Srv     *DepartmentService
}

func NewDepartmentModule(infra *storage.Infra) *DepartmentModule {
	repo := NewDepartmentRepository(infra.MySQL)
	srv := NewDepartmentService(repo)

	return &DepartmentModule{Handler: &DepartmentHandler{srv: srv}, Srv: srv}
}
