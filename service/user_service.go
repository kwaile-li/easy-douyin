package service

import (
	"douyin/dao"
)

type UserService interface {

	// GetUserBasicInfoById 根据Id获取用户的用户名和密码
	GetUserBasicInfoById(id int64) dao.UserBasicInfo

	// GetUserBasicInfoByName 根据用户名获取用户的用户名和密码
	GetUserBasicInfoByName(name string) dao.UserBasicInfo

	// GetUserLoginInfoById 根据用户id获取用户的详细信息（未登录）
	GetUserLoginInfoById(id int64) (dao.User, error)

	// GetUserLoginInfoByIdWithCurId 根据用户id获取用户的详细信息 (登录)
	GetUserLoginInfoByIdWithCurId(id int64, curId int64) (dao.User, error)

	// InsertUser 添加一个用户
	InsertUser(user *dao.UserBasicInfo) bool
}
