package constant

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "invalid request parameters",
	FAIL_ADD_DATA : "failed to save data",

	// user-related
	ERROR_EXIST_USER : "username already exists",
	ERROR_NOT_EXIST_USER : "user does not exist",
	ERROR_PASS_USER : "incorrect password",
	ERROR_CAPTCHA_USER : "invalid captcha",
	FAIL_LOGOUT_USER : "logout failed",
	// token-related
	ERROR_AUTH_CHECK_TOKEN_FAIL : "token authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "token expired",
	ERROR_AUTH_TOKEN : "failed to generate token",
	ERROR_AUTH : "invalid token",
	ERROR_AUTH_CHECK_FAIL : "permission denied, contact administrator",

	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "failed to save image",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "failed to validate image",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "invalid image format or size",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
