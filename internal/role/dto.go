package role

import "go-learn/internal/common/utils"

type RoleCreateRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Code        string `json:"code" validate:"required,max=255"`
	Description string `json:"description" validate:"max=255"`
}

type RoleInfoResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func toDetailResponse(role *Role) *RoleInfoResponse {
	if role == nil {
		return nil
	}
	return &RoleInfoResponse{
		ID:          role.ID,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		CreatedAt:   utils.FmtDateTime(role.CreatedAt),
		UpdatedAt:   utils.FmtDateTime(role.UpdatedAt),
	}
}
