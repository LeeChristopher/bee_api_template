package services

import (
	"bee_api_template/models/user"
	"bee_api_template/packages"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

/**
 *登录
 */
func Login(username string, password string) (userInfo *user.User, err error) {
	userInfo = &user.User{}
	err = packages.Db.Table(user.GetTableName()).Select(user.GetLoginField()).
		Where("name = ?", username).
		Find(userInfo).Error
	if err != nil {
		return userInfo, errors.New("账户信息不存在！")
	}
	if strings.Compare(password, userInfo.Password) != 0 {
		return userInfo, errors.New("用户名或密码错误！")
	}

	authUserInfo := packages.AuthUserInfo{
		UserId:   userInfo.Id,
		UserName: userInfo.Name,
	}
	result, token := packages.IssueAuthToken(authUserInfo)
	if !result {
		return userInfo, errors.New("登录失败！")
	}
	userInfo.AccessToken = token

	tokenKey := packages.GetLoginKey(userInfo.Id)
	authByte, err := json.Marshal(userInfo)
	if err != nil {
		return userInfo, errors.New("登录失败！")
	}
	err = packages.Redis.Set(tokenKey, string(authByte), time.Second*packages.TokenExpire).Err()
	if err != nil {
		return userInfo, errors.New("登录失败！")
	}

	return
}
