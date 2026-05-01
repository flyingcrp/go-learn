package department

import "go-learn/internal/common/utils"

type DepartmentCreateRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"max=500"`
}

type DepartmentInfoResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func toDetailResponse(dep *Department) *DepartmentInfoResponse {
	if dep == nil {
		return nil
	}
	return &DepartmentInfoResponse{
		ID:          dep.ID,
		Name:        dep.Name,
		Description: dep.Description,
		CreatedAt:   utils.FmtDateTime(dep.CreatedAt),
		UpdatedAt:   utils.FmtDateTime(dep.UpdatedAt),
	}
}
