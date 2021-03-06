package config

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 601
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 602
	ERROR_AUTH_TOKEN               = 603
	ERROR_AUTH                     = 604

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 888
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 889
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 890
)

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
