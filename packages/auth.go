package packages

import (
	"bee_api_template/models/user"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	AppKey      = ""
	TokenExpire = 86400
)

/*
 * token中存储的用户信息
 */
type AuthUserInfo struct {
	UserId   uint64 `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type AuthTokenClaims struct {
	AuthUserInfo
	jwt.StandardClaims
}

/*
 * 验证认证Token
 */
func VerifyAuthToken(tokenString string) (userId uint64, err error) {
	tokenObj, err := jwt.ParseWithClaims(tokenString, &AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(AppKey), nil
	})
	if err != nil {
		return 0, errors.New("invalid token")
	}
	if tokenObj == nil || !tokenObj.Valid {
		return 0, errors.New("invalid token")
	}
	authUserInfo, ok := tokenObj.Claims.(*AuthTokenClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	redisCache, err := Redis.Get(GetLoginKey(authUserInfo.UserId)).Bytes()
	if err != nil {
		return 0, errors.New("invalid token")
	}
	if len(redisCache) == 0 {
		return 0, errors.New("invalid token")
	}
	userInfo := user.User{}
	err = json.Unmarshal(redisCache, &userInfo)
	if err != nil {
		return 0, errors.New("invalid token")
	}
	if userInfo.AccessToken != tokenString {
		return 0, errors.New("invalid token")
	}

	return userInfo.Id, nil
}

/*
 * 签发认证Token
 */
func IssueAuthToken(userInfo AuthUserInfo) (result bool, tokenString string) {
	claim := AuthTokenClaims{
		AuthUserInfo: userInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: GetNow().Add(time.Second * TokenExpire).Unix(),
			Issuer:    userInfo.UserName,
			IssuedAt:  GetNow().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(AppKey))
	if err != nil {
		return false, ""
	}

	return true, tokenString
}
