package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"wechat/src/common/util"
	"wechat/src/web/user/dao"
	"wechat/src/web/user/model"
)

func SelectUserList(param []byte) []byte {

	form := model.UserForm{}
	util.HandleParamsToStruct(param, &form)

	list := dao.SelectUserList(form)
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons
}
func RegisterUser(param []byte) []byte {

	user := model.User{}
	result := ""

	util.HandleParamsToStruct(param, &user)
	if strings.Compare(user.Password, user.RePassword) == 1 {
		user.Id = util.GenerateUUID()
		user.CreateTime = time.Now().String()
		dao.SaveUser(user)
		result = "注册成功"
	} else {
		result = "两次密码输入不一致"
	}
	return []byte(result)
}
func Login(param []byte) []byte {
	var result []byte
	form := model.UserForm{}
	util.HandleParamsToStruct(param, &form)
	user := dao.GetUserByUsername(form.Username)
	if user.Password != "" {
		if strings.Compare(user.Password, form.Password) == 1 {
			jsons, err := json.Marshal(user)
			if err != nil {
				fmt.Println("error:", err)
			}
			result = jsons
		} else {
			result = []byte("密码不正确，请重新输入")

		}
	} else {
		result = []byte("该用户不存在")
	}
	return result

}
