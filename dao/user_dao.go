package dao

import (
	"log"
	"time"
)

type UserBasicInfo struct {
	Id        int64
	Name      string
	Password  string
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}

type User struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

func (user UserBasicInfo) TableName() string {
	return "UserBasicInfo"
}

// 插入用户
func InsertUser(user *UserBasicInfo) bool {
	if err := Db.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// 根据用户名查取用户信息
func GetUserBasicInfoByName(name string) (UserBasicInfo, error) {
	user := UserBasicInfo{}
	if err := Db.Where("name = ?", name).First(&user).Error; err != nil {
		log.Println("获取用户信息读库失败", err.Error())
		return user, err
	}
	return user, nil
}

// 根据用户id查取用户信息
func GetUserBasicInfoById(id int64) (UserBasicInfo, error) {
	user := UserBasicInfo{}
	if err := Db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}
