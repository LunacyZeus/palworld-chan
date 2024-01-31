package controllers

import (
	_ "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"palworld-chan/internal/service/api/config"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/api/pkg/resp"
	"time"
)

func Login(c *fiber.Ctx) error { //服务器状态 获取
	// 解析 JSON 数据到结构体
	var user models.UserInput
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// 处理用户数据
	// 你可以在这里进行登录验证等操作
	if !config.Auth().Check(user.Username, user.Password) {
		//{"code":300,"result":null,"message":"帐号或密码不正确","type":"error"}
		res := models.Response{
			Code:    300,
			Result:  nil,
			Message: "帐号或密码不正确",
			Type:    "error",
		}
		return c.JSON(res)

	}
	//jwt
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user.Username,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// 返回响应
	result := models.LoginResult{
		UserID:   0,
		Username: user.Username,
		Token:    t,
		Realname: user.Username,
		Sign:     "hello world",
	}
	//{"code":200,"result":{"userId":2,"username":"hoho","token":"fakeToken2","realname":"test user","sign":"这个家伙很懒，什么都没有写~"},"message":"ok","type":"success"}
	res := models.Response{
		Code:    200,
		Result:  result,
		Message: "ok",
		Type:    "success",
	}

	return c.JSON(res)
}

func GetUserInfo(c *fiber.Ctx) error {
	// 返回响应
	result := models.UserInfoResult{
		UserID:   0,
		Username: "",
		PassWord: "",
		NickName: "",
		RealName: "",
		Avatar:   "",
		Cover:    "",
		Sign:     "",
		Industry: 0,
		Gender:   0,
		Thone:    "",
		Token:    "",
	}

	res := resp.JsonResp(result)

	return c.Status(fiber.StatusOK).JSON(res)
}
