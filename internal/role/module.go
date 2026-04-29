package role

import (
	"go-learn/internal/common/storage"
	"time"
)

type Role struct {
	ID          string     `gorm:"size:36;primaryKey;type:varchar(36)"`
	Name        string     `gorm:"size:255;not null"`
	Code        string     `gorm:"size:255;not null;unique"`
	Description string     `gorm:"size:255"`
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
}

func (Role) TableName() string {
	return "role"
}

type RoleModule struct {
	Handler *RoleHandler
	Utils   *Utils
}

func NewRoleModule(infra *storage.Infra) *RoleModule {
	repo := NewRoleRepository(infra.MySQL)
	srv := NewRoleService(repo)
	return &RoleModule{Handler: NewRoleHandler(srv), Utils: NewUtils(repo)}
}
