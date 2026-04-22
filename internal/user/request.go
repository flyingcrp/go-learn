package user

type CreateUserReq struct {
	Name  string `json:"name" binding:"required,min=2,max=20"`
	Email string `json:"email" binding:"required,email"`
}
