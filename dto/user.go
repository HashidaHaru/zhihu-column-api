package dto

// SignupDto 注册
type SignupDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginDto 登陆
type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
