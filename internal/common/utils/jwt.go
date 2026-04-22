package utils

type JWTClaims struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyID string `json:"company_id"`
}
