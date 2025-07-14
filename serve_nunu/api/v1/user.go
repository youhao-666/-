package v1

type RegisterRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required" `
}
