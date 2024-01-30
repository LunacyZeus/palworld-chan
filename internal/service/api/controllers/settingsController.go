package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/api/pkg/custom"
	"palworld-chan/internal/service/dao"
	utils "palworld-chan/pkg/utility/utils"
	"strings"
)

func UpdateServerSetting(c *fiber.Ctx) error { //发送服务器广播
	// 解析 JSON 数据到结构体
	var updateServerSettingInput models.UpdateServerSettingInput
	if err := c.BodyParser(&updateServerSettingInput); err != nil {
		return err
	}

	// Validation
	if errs := custom.CustomValidator().Validate(updateServerSettingInput); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}

	//解析成功后 转换到持久层
	// 转换为 JSON 字符串
	serverSettingJson, err := utils.ToJSONString(updateServerSettingInput)
	if err != nil {
		return err
	}

	_ = dao.Set(consts.BUCKET, "serverSetting", serverSettingJson)

	res := models.Response{
		Code:    200,
		Result:  nil,
		Message: "updated",
		Type:    "success",
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func GetServerSetting(c *fiber.Ctx) error { //发送服务器广播
	//从持久层解析到结构体
	serverSettingJson, err := dao.Get(consts.BUCKET, "serverSetting")

	if err != nil {
		return err
	}

	// 创建 UpdateServerSettingInput 实例
	var updateServerSettingInput models.UpdateServerSettingInput

	// 解析 JSON 字符串到结构体
	err = utils.FromJSONString(serverSettingJson, &updateServerSettingInput)
	if err != nil {
		return err
	}

	res := models.Response{
		Code:    200,
		Result:  updateServerSettingInput,
		Message: "",
		Type:    "success",
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
