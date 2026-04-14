package token

import (
	"TikTok/biz/model/user"
	"TikTok/dal/mysql"
	"TikTok/utils"
	"context"

	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

var JwtAuth *jwt.HertzJWTMiddleware

type UserClaims struct {
	Username string `json:"username"`
	UserId   string `json:"user_id"`
}

func JwtAuthenticator(c context.Context, req *app.RequestContext) (interface{}, error) {
	var loginVals user.LoginRequest
	if err := req.BindAndValidate(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	var use user.User
	mysql.Db.Where("username = ?", loginVals.Username).First(&use)
	if utils.CompareHashAndPassword(loginVals.Password, use.Password) {
		return &UserClaims{
			Username: loginVals.Username,
			UserId:   use.Id,
		}, nil
	} else {
		return nil, jwt.ErrFailedAuthentication
	}
}
func JwtPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*UserClaims); ok {
		return jwt.MapClaims{
			"IdentityKey": v.UserId,
			"username":    v.Username,
		}
	}
	return jwt.MapClaims{}
}

func JwtUnauthorized(c context.Context, req *app.RequestContext, code int, message string) {
	req.JSON(code, map[string]interface{}{
		"code": code,
		"msg":  message,
	})
}

func NewAuthMiddleware() {
	var err error
	JwtAuth, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            "TikTok",
		Key:              []byte(Jwtconfig.Jwtsecret),
		SigningAlgorithm: "HS256",
		Timeout:          time.Duration(Jwtconfig.Accesstokenexpire) * time.Second,
		MaxRefresh:       time.Duration(Jwtconfig.Refreshtokenexpire) * time.Second,
		IdentityKey:      Jwtconfig.Identitykey,
		TokenHeadName:    "Bearer",
		Unauthorized:     JwtUnauthorized,
		TokenLookup:      "header:Authorization",
	})
	if err != nil {
		panic("JWT Error:" + err.Error())
	}

}

func JWT() app.HandlerFunc {
	return JwtAuth.MiddlewareFunc()
}
