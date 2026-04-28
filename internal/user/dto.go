package user

type UserRegisterReq struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID string `json:"department_id" binding:"required"`
	RoleID       string `json:"role_id" binding:"required"`
}
type UserRegisterResp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUpdateReq struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
type UserUpdateResp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
