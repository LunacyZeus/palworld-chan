package resp

import "palworld-chan/internal/service/api/models"

func JsonResp(result any) models.Response {
	resp := models.Response{
		Code:    200,
		Result:  result,
		Message: "ok",
		Type:    "success",
	}
	return resp
}
