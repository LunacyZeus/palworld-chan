package models

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Realname string `json:"realname"`
	Sign     string `json:"sign"`
}

type UserInfoResult struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	PassWord string `json:"password"`
	NickName string `json:"nickname"`
	RealName string `json:"realname"`
	Avatar   string `json:"avatar"`
	Cover    string `json:"Cover"`
	Sign     string `json:"sign"`
	Industry int    `json:"industry"`
	Gender   int    `json:"gender"`
	Thone    string `json:"phone"`
	Token    string `json:"token"`
}
