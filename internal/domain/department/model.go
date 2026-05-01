package department

import "time"

type Department struct {
	ID          string     `gorm:"size:36;primaryKey;type:varchar(36)"`
	Name        string     `gorm:"size:255;not null"`
	Description string     `gorm:"size:255;type:varchar(500)"`
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
}

func (Department) TableName() string {
	return "department"
}
