package user

import (
	"context"
	"go-learn/internal/domain/department"
	"go-learn/internal/domain/role"
	"go-learn/internal/infra/event"
	"testing"
	"time"

	"github.com/google/uuid"
)

// ============================================================
// Mock 实现：为了 benchmark 准备，不依赖真实 DB / 外部服务
// ============================================================

// stubUserRepo 实现 UserRepo 接口，所有方法返回预设值
type stubUserRepo struct {
	// 可配置的返回值，方便不同测试场景
	createErr        error
	existsByEmail    bool
	existsByEmailErr error
	findByIDUser     *User
	findByIDErr      error
	listUsers        []User
	listTotal        int64
	listErr          error
}

func (m *stubUserRepo) Create(ctx context.Context, user *User) error {
	return m.createErr
}

func (m *stubUserRepo) Update(ctx context.Context, id string, user *User) error {
	return nil
}

func (m *stubUserRepo) FindByID(ctx context.Context, id string) (*User, error) {
	return m.findByIDUser, m.findByIDErr
}

func (m *stubUserRepo) Login(ctx context.Context, name, email string) (*User, error) {
	return m.findByIDUser, nil
}

func (m *stubUserRepo) List(ctx context.Context, params *UserListReq) ([]User, int64, error) {
	return m.listUsers, m.listTotal, m.listErr
}

func (m *stubUserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return m.existsByEmail, m.existsByEmailErr
}

// stubDepartmentChecker 实现 departmentChecker 接口
type stubDepartmentChecker struct {
	dep *department.Department
	err error
}

func (m *stubDepartmentChecker) CheckID(ctx context.Context, id string) (*department.Department, error) {
	return m.dep, m.err
}

// stubRoleChecker 实现 roleChecker 接口
type stubRoleChecker struct {
	role *role.Role
	err  error
}

func (m *stubRoleChecker) CheckID(ctx context.Context, id string) (*role.Role, error) {
	return m.role, m.err
}

// ============================================================
// 预分配的测试数据（避免 benchmark 重复分配干扰结果）
// ============================================================

var (
	now = time.Now()

	benchUser = &User{
		ID:           uuid.New().String(),
		Name:         "bench",
		Email:        "bench@example.com",
		DepartmentID: uuid.New().String(),
		RoleID:       uuid.New().String(),
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	benchDep  = &department.Department{ID: benchUser.DepartmentID, Name: "Engineering"}
	benchRole = &role.Role{ID: benchUser.RoleID, Name: "Developer", Code: "dev"}

	benchListUsers = []User{*benchUser, *benchUser, *benchUser}
)

// newBenchService 构造一个用于 benchmark 的 UserService，所有依赖都已 mock
func newBenchService() *UserService {
	repo := &stubUserRepo{
		existsByEmail: false,
		listUsers:     benchListUsers,
		listTotal:     3,
	}
	depChecker := &stubDepartmentChecker{dep: benchDep}
	roleChecker := &stubRoleChecker{role: benchRole}
	bus := event.NewBus()
	return NewUserService(repo, depChecker, roleChecker, bus, "bench-secret")
}

// ============================================================
// Benchmarks
// ============================================================

// TODO: 完成 BenchmarkRegister — 测量 Register 的 happy path
func BenchmarkRegister(b *testing.B) {
	// 1. 用 newBenchService() 构造被测对象
	svc := newBenchService()
	// 2. 构造 UserRegisterReq
	req := UserRegisterReq{
		Name:         "bench",
		Email:        "bench@example.com",
		DepartmentID: benchDep.ID,
		RoleID:       benchRole.ID,
	}
	// 3. b.ResetTimer() 后进入 b.N 循环
	b.ResetTimer()
	// 4. 循环内只调用 svc.Register
	for i := 0; i < b.N; i++ {
		_, _ = svc.Register(context.Background(), &req)
	}
}
