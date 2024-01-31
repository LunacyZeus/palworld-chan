package config

import (
	"sync"
)

type WebAuth struct {
	UserName string
	PassWord string
}

func (c *WebAuth) Set(UserName, PassWord string) {
	c.UserName = UserName
	c.PassWord = PassWord
}

func (c *WebAuth) Check(UserName, PassWord string) bool {
	// 检查用户名和密码
	//log.Println(UserName, c.UserName, PassWord, c.PassWord)
	if UserName == c.UserName && PassWord == c.PassWord {
		return true
	}
	return false
}

var instance *WebAuth
var once sync.Once

func Auth() *WebAuth {
	once.Do(func() {
		webAuth := &WebAuth{
			UserName: "pal",
			PassWord: "123",
		}
		instance = webAuth
	})
	return instance
}
