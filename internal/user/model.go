package user

type User struct {
	ID           string `gorm:"size:36;primaryKey;type:varchar(36)"`
	Name         string `gorm:"size:255;not null"`
	Email        string `gorm:"size:255;not null;unique"`
	DepartmentID string `gorm:"size:36;not null"`
	RoleID       string `gorm:"size:36;not null"`
	Token        string `gorm:"size:255;"`
}

func (User) TableName() string {
	return "user"
}
