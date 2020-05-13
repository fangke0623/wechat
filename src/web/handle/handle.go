package handle

import (
	"encoding/json"
	"log"
	"net/http"
	"wechat/src/common/exception"
	"wechat/src/web/discuss"
	"wechat/src/web/discussMan"
	"wechat/src/web/discussMsg"
	"wechat/src/web/user"
)

func ResponseHandle(writer http.ResponseWriter, request *http.Request) {

	param, err := handleParams(request)
	if err != nil {
		log.Print("form转换json异常：", err)
	}
	var result interface{}
	var e exception.Error
	switch request.URL.Path {
	//user
	case "/user/list":
		result, e = user.FindUserList(param)
		break
	case "/user/detail":
		result, e = user.DetailUser(param)
		break
	case "/user/register":
		result, e = user.RegisterUser(param)
		break
	case "/user/login":
		result, e = user.Login(param)
		break
	case "/user/edit":
		result, e = user.Edit(param)
		break

	//discuss
	case "/discuss/add":
		result, e = discuss.AddDiscuss(param)
		break
	case "/discuss/update":
		result, e = discuss.UpdateDiscuss(param)
		break
	case "/discuss/delete":
		result, e = discuss.DeleteDiscuss(param)
		break
	case "/discuss/list":
		result, e = discuss.FindDiscussList(param)
		break
	//discussMan

	case "/discussMan/list":
		result, e = discussMan.FindDiscussManList(param)
		break
	case "/discussMan/add":
		result, e = discussMan.AddDiscussMan(param)
		break
	case "/discussMan/update":
		result, e = discussMan.AddDiscussMan(param)
		break

	//discussMsg
	case "/discussMsg/list":
		result, e = discussMsg.FindDiscussMsgList(param)
		break
	case "/discussMsg/add":
		result, e = discussMsg.AddDiscussMsg(param)
		break
	}
	writerJson(writer, result, e)
}
func handleParams(request *http.Request) ([]byte, error) {
	var param []byte
	var err error
	_ = request.ParseForm()
	if request.Method == "POST" {
		param, err = json.Marshal(request.PostForm)
	}
	if request.Method == "GET" {
		param, err = json.Marshal(request.Form)
	}
	return param, err
}

func writerJson(writer http.ResponseWriter, param interface{}, exception exception.Error) {
	resultMap := make(map[string]interface{})
	resultMap["data"] = param
	resultMap["errorCode"] = exception.ErrorCode
	resultMap["errorMsg"] = exception.ErrorMsg
	result, _ := json.Marshal(resultMap)
	_, err := writer.Write(result)
	if err != nil {
		log.Println(err)
	}
}
