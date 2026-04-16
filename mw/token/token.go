package token

import (
	"net/http"
	"time"

	myredis "TikTok/dal/redis"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
)

func SaveRefreshToken(userId string, refreshToken string, expiiration time.Duration) error {
	key := "refresh_token:" + userId
	return myredis.Rdb.Set(context.Background(), key, refreshToken, expiiration).Err()
} //保存刷新令牌到redis，并设置过期时间，为什么这里可以自动刷新呢？因为redis采用的是key-value的存储方式，每次访问这个key都会重置过期时间，所以只要用户在使用过程中不断访问，就不会过期，只有当用户长时间不使用时才会过期
func DeleteRefreshToken(userId string) error {
	key := "refresh_token:" + userId
	return myredis.Rdb.Del(context.Background(), key).Err()
} //删除redis中保存的刷新令牌

func CheckRefreshToken(userId string, refreshToken string) (bool, error) {
	key := "refresh_token:" + userId
	storedToken, err := myredis.Rdb.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil // 没有找到刷新令牌
		}
		return false, err // 其他错误
	}
	return storedToken == refreshToken, nil
} //检查是否存在有效的刷新令牌

func GenerateToken(userId string, username string) (string, string, error) {
	accessclaims := jwt.MapClaims{
		"userid":   userId,
		"username": username,
		"type":     "access",
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Minute * 24)), // 设置过期时间为24f小时
	}
	freshclaims := jwt.MapClaims{
		"userid":   userId,
		"username": username,
		"type":     "refresh",
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 设置过期时间为24小时
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessclaims)
	freshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, freshclaims)

	accessTokenString, err := accessToken.SignedString([]byte(Jwtconfig.Jwtsecret))
	if err != nil {
		return "", "", err
	}
	freshTokenString, err := freshToken.SignedString([]byte(Jwtconfig.Jwtsecret))
	if err != nil {
		return "", "", err
	}
	err = SaveRefreshToken(userId, freshTokenString, time.Hour*24)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, freshTokenString, nil
} //生成token字符串，并保存刷新令牌到redis
func ParseToken(tokenString string) *UserClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Jwtconfig.Jwtsecret), nil
	})
	if err != nil {
		return nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &UserClaims{
			UserId:   claims["userid"].(string),
			Username: claims["username"].(string),
		}
	}
	return nil
} // 解析token字符串，返回用户信息

func RefreshToken(ctx context.Context, req *app.RequestContext) {
	cookie := req.Cookie("refresh_token")

	if len(cookie) == 0 {
		req.JSON(401, map[string]interface{}{
			"code": 401,
			"msg":  "未找到刷新令牌",
		})
		return
	}
	oldReshToken := string(cookie)
	user := ParseToken(oldReshToken)
	if user == nil {
		req.JSON(401, map[string]interface{}{
			"code": 401,
			"msg":  "刷新令牌无效",
		})
		return
	}
	userexists, err := CheckRefreshToken(user.UserId, oldReshToken)
	if err != nil {
		req.JSON(500, map[string]interface{}{
			"code": 500,
			"msg":  "检查刷新令牌时发生错误",
		})
		return
	}
	if !userexists {
		req.JSON(401, map[string]interface{}{
			"code": 401,
			"msg":  "刷新令牌无效",
		})
		return
	}
	accessclaims := jwt.MapClaims{
		"userid":   user.UserId,
		"username": user.Username,
		"type":     "access",
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Minute * 24)), // 设置过期时间为24f小时
	}
	freshclaims := jwt.MapClaims{
		"userid":   user.UserId,
		"username": user.Username,
		"type":     "refresh",
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 设置过期时间为24小时
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessclaims)
	freshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, freshclaims)

	accessTokenString, err := accessToken.SignedString([]byte(Jwtconfig.Jwtsecret))
	if err != nil {
		req.JSON(500, map[string]interface{}{
			"code": 500,
			"msg":  "生成访问令牌时发生错误",
		})
		return
	}
	freshTokenString, err := freshToken.SignedString([]byte(Jwtconfig.Jwtsecret))
	if err != nil {
		req.JSON(500, map[string]interface{}{
			"code": 500,
			"msg":  "生成刷新令牌时发生错误",
		})
		return
	}
	err = DeleteRefreshToken(user.UserId)
	if err != nil {
		req.JSON(500, map[string]interface{}{
			"code": 500,
			"msg":  "删除旧刷新令牌时发生错误",
		})
		return
	}

	err = SaveRefreshToken(user.UserId, freshTokenString, time.Hour*24)
	if err != nil {
		req.JSON(500, map[string]interface{}{
			"code": 500,
			"msg":  "保存刷新令牌时发生错误",
		})
		return
	}
	req.SetCookie("refresh_token", freshTokenString, 3600*24, "/", "", protocol.CookieSameSite(http.SameSiteLaxMode), true, true)

	req.JSON(200, map[string]interface{}{
		"code":         200,
		"msg":          "令牌刷新成功",
		"access_token": accessTokenString,
	})
} //刷新了令牌
