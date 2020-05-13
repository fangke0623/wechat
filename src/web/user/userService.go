package user

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"wechat/src/common/exception"
	"wechat/src/common/util"
)

func FindUserList(param []byte) (interface{}, exception.Error) {

	e := exception.Error{}
	form := Form{}
	util.HandleParamsToStruct(param, &form)

	list := SelectUserList(form)
	return list, e
}
func DetailUser(param []byte) ([]byte, exception.Error) {

	form := Form{}
	e := exception.Error{}

	util.HandleParamsToStruct(param, &form)

	user := GetUserById(form.Id)
	jsons, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons, e
}
func RegisterUser(param []byte) ([]byte, exception.Error) {

	user := User{}
	result := ""
	e := exception.Error{}
	util.HandleParamsToStruct(param, &user)
	dbUser := GetUserByUsername(user.UserName)
	if dbUser.UserName != "" {
		e = exception.UserNameIsExist
		return []byte(result), e
	}
	if strings.Compare(user.Password, user.RePassword) == 1 {
		user.Id = util.GenerateUUID()
		user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		SaveUser(user)
		e.ErrorMsg = "注册成功"
	} else {
		e = exception.PassWordIsInconsistent
		return []byte(result), e
	}
	return []byte(result), e
}
func Login(param []byte) ([]byte, exception.Error) {
	e := exception.Error{}
	form := Form{}
	jsons := []byte("")
	util.HandleParamsToStruct(param, &form)
	user := GetUserByUsername(form.Username)
	if user.Password != "" {
		if strings.Compare(user.Password, form.Password) == 1 {
			jsons, err := json.Marshal(user)
			if err != nil {
				fmt.Println("error:", err)
			}
			return jsons, e
		} else {
			e = exception.PassWordIsWrong
		}
	} else {
		e = exception.UserNotExist
	}
	return jsons, e

}
func Edit(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	user := User{}
	util.HandleParamsToStruct(param, &user)
	dbUser := GetUserById(user.Id)
	if dbUser.UserName == "" {
		e = exception.UserNotExist
	}

	user.Id = util.GenerateUUID()
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	UpdateUserById(user)
	e.ErrorMsg = "修改成功"

	return "", e
}