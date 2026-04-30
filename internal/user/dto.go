package user

import (
	"go-learn/internal/common/utils"
)

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

type UserLoginReq struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
type UserLoginResp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func toLoginResp(user *User, token string) *UserLoginResp {
	if user == nil {
		return nil
	}
	return &UserLoginResp{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}

type UserDetailResp struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	DepartmentID string `json:"department_id"`
	RoleID       string `json:"role_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func toUserDetailResp(user *User) *UserDetailResp {
	if user == nil {
		return nil
	}
	return &UserDetailResp{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		DepartmentID: user.DepartmentID,
		RoleID:       user.RoleID,
		CreatedAt:    utils.FmtDateTime(user.CreatedAt),
		UpdatedAt:    utils.FmtDateTime(user.UpdatedAt),
	}
}

type UserListReq struct {
	utils.Pagination
	Name string `form:"name"`
}
type UserListResp struct {
	List  []UserDetailResp `json:"list"`
	Total int64            `json:"total"`
}
