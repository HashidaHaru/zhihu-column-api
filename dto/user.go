package dto

// SignupDto 注册
type SignupDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// LoginDto 登陆
type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
