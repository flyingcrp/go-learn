package user

import "time"

type User struct {
	ID           string     `gorm:"size:36;primaryKey;type:varchar(36)"`
	Name         string     `gorm:"size:255;not null"`
	Email        string     `gorm:"size:255;not null;unique"`
	DepartmentID string     `gorm:"size:36;not null"`
	RoleID       string     `gorm:"size:36;not null"`
	Token        string     `gorm:"size:255;"`
	CreatedAt    *time.Time `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}
