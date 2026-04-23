package user

type CreateUserReq struct {
	ID    uint64 `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required,min=2,max=20"`
	Email string `json:"email" binding:"required,email"`
}

type User struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"size:20;not null"`
	Email string `json:"email" gorm:"size:100;not null;uniqueIndex"`
}
