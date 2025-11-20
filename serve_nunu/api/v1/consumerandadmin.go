package v1

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

type UpdateReq struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
