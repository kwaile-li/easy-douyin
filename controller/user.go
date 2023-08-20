package controller

import (
	"douyin/dao"
	"douyin/service"
	"douyin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User dao.User `json:"user"`
}

/*
username string 必需
注册用户名，最长32个字符

password string 必需
密码，最长32个字符
*/
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	usi := service.GetUserServiceInstance()
	user := usi.GetUserBasicInfoByName(username)
	if username == user.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := dao.UserBasicInfo{
			Name:     username,
			Password: service.EnCoder(password),
		}
		if usi.InsertUser(&newUser) != true { // 插入新用户
			fmt.Println("Insert Fail")
		}
		// 得到用户id
		user := usi.GetUserBasicInfoByName(username)
		userId := user.Id
		token := util.GenerateToken(userId, username) // jwt获取token

		//log.Println("注册时返回的token", token)
		//log.Println("注册返回的id: ", user.Id)
		// 返回token、userid和其他信息
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Register Success"},
			UserId:   user.Id,
			Token:    token,
		})
	}
}
