package department

import "go-learn/internal/common/storage"

func NewDepartmentModule(infra *storage.Infra) *DepartmentHandler {
	repo := NewDepartmentRepository(infra.MySQL)
	srv := NewDepartmentService(repo)
	return &DepartmentHandler{srv: srv}
}
