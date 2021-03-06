package exception

type Error struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

var (
	// common
	Success   = Error{00, "请求成功"}
	ParamMiss = Error{01, "必要参数为空"}

	//user
	PassWordIsWrong        = Error{10, "密码错误"}
	PassWordIsInconsistent = Error{11, "两次输入密码不一致"}
	UserNotExist           = Error{12, "用户不存在"}
	UserNameIsExist        = Error{13, "用户名已存在"}

	//discuss
	DiscussNotExist = Error{22, "聊天室不存在"}

	//discussMan
	DiscussManNotExist = Error{32, "人员不存在"}

	//discussMsg
)

func ThrowException(error Error) (interface{}, Error) {

	return "", error
}
