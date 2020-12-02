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

// UserModify 修改用户信息
type UserModify struct {
	Email    string `json:"email" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}
