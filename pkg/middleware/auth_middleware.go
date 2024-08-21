package middleware

import (
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.corp.yxkj.com/guava/common/config"
)

var TokenHeadName = "Bearer" // Token 认证方式

// InitAuth
// 初始化 JWT 中间件
func InitAuth(ai AuthInterface) (*jwt.GinJWTMiddleware, error) {
	Authdleware, err := jwt.New(&jwt.GinJWTMiddleware{
		DisabledAbort:   true,                                               //禁止此三方包内部Abort()
		Realm:           ai.GetJwtRealm(),                                   // jwt 标识
		Key:             []byte(ai.GetJwtKey()),                             // 服务端密钥
		Timeout:         time.Hour * time.Duration(ai.GetJwtTimeOut()),      // token 过期时间
		MaxRefresh:      time.Hour * time.Duration(ai.GetJwtMaxRefresh()),   // token 最大刷新时间(RefreshToken 过期时间 = Timeout+MaxRefresh)
		PayloadFunc:     ai.PayloadFunc,                                     // 有效载荷处理
		IdentityHandler: ai.IdentityHandler,                                 // 解析 Claims
		Authenticator:   ai.Login,                                           // 校验 token 的正确性, 处理登录逻辑
		Authorizator:    ai.Authorizator,                                    // 用户登录校验成功处理
		Unauthorized:    ai.Unauthorized,                                    // 用户登录校验失败处理
		LoginResponse:   ai.LoginResponse,                                   // 登录成功后的响应
		LogoutResponse:  ai.LogoutResponse,                                  // 登出后的响应
		RefreshResponse: ai.RefreshResponse,                                 // 刷新 token 后的响应
		TokenLookup:     "header: Authorization, query: token, cookie: jwt", // 自动在这几个地方寻找请求中的 token
		TokenHeadName:   TokenHeadName,                                      // header 名称
		TimeFunc:        time.Now,
	})
	return Authdleware, err
}

type AuthIntf interface {
	Login(c *gin.Context) (interface{}, error)
	Authorizator(data interface{}, c *gin.Context) bool
	Unauthorized(ctx *gin.Context, code int, message string)
	LoginResponse(ctx *gin.Context, _ int, token string, expires time.Time)
	LogoutResponse(ctx *gin.Context, _ int)
	RefreshResponse(ctx *gin.Context, _ int, token string, expires time.Time)
	PayloadFunc(data interface{}) string
}

type Auth struct {
	auth AuthIntf
}

func NewAuth(a AuthIntf) *jwt.GinJWTMiddleware {
	auth := new(Auth)
	auth.auth = a
	jwt, err := InitAuth(auth)
	if err != nil {
		panic(fmt.Sprintf("初始化鉴权中间件失败，失败原因：%s", err.Error()))
	}
	return jwt
}

// PayloadFunc
// 有效载荷处理
func (auth *Auth) PayloadFunc(data interface{}) jwt.MapClaims {
	mapClaims := make(jwt.MapClaims)
	v, ok := data.(map[string]interface{})
	if ok {
		id := auth.auth.PayloadFunc(v["info"])
		mapClaims[jwt.IdentityKey] = id
		mapClaims["info"] = v["info"]
	}
	return mapClaims
}

// IdentityHandler
// 解析Claims
func (auth *Auth) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	// 此处返回值类型 map[string]interface{}
	// 与 payloadFunc 和 authorizator 的 data 类型必须一致, 否则会导致授权失败还不容易找到原因
	mapData := make(map[string]interface{})
	mapData["IdentityKey"] = claims[jwt.IdentityKey]
	mapData["info"] = claims["info"]
	return mapData
}

// Login godoc
//
//	@tags			基本接口
//	@Summary		用户登录
//	@Produce		json
//	@Description	```
//	@Description	用户登录
//	@Description	```
//	@Param			data	body		repository.Login	true	"登录信息"
//	@Success		200		{object}	controller.Full
//	@Failure		500		{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/login [post]
func (auth *Auth) Login(c *gin.Context) (interface{}, error) {
	return auth.auth.Login(c)
}

// Authorizator
// 用户登录校验成功处理
func (auth *Auth) Authorizator(data interface{}, c *gin.Context) bool {
	return auth.auth.Authorizator(data, c)
}

// Unauthorized
// 用户登录校验失败处理
func (auth *Auth) Unauthorized(ctx *gin.Context, code int, message string) {
	auth.auth.Unauthorized(ctx, code, message)
}

// LoginResponse
// 登录成功后的响应
func (auth *Auth) LoginResponse(ctx *gin.Context, i int, token string, expires time.Time) {
	auth.auth.LoginResponse(ctx, i, token, expires)
}

// LogoutResponse godoc
//
//	@tags			基本接口
//	@Summary		用户登出
//	@Description	用户登出时，调用此接口
//	@Produce		json
//	@Success		200	{object}	controller.Base
//	@Failure		500	{object}	controller.Base	"错误返回内容"
//	@Router			/api/v1/logout [get]
//	@Security		ApiKeyAuth
func (auth *Auth) LogoutResponse(ctx *gin.Context, i int) {
	auth.auth.LogoutResponse(ctx, i)
}

// RefreshResponse godoc
//
//	@tags			基本接口
//	@Summary		刷新token
//	@Description	当token即将获取或者过期时刷新token
//	@Produce		json
//	@Success		200	{object}	controller.Full	"code:200 成功"
//	@Failure		500	{object}	controller.Base					"错误返回内容"
//	@Router			/api/v1/refreshToken [get]
//	@Security		ApiKeyAuth
func (auth *Auth) RefreshResponse(ctx *gin.Context, i int, token string, expires time.Time) {
	auth.auth.RefreshResponse(ctx, i, token, expires)
}

// GetJwtRealm
// 获取 jwt标识
func (auth *Auth) GetJwtRealm() string {
	return config.GetParam("JWT-REALM", "042f7a4b82bb4c48a9cb3082a47818532765c0cc").String()
}

// GetJwtKey
// jwt 秘钥
func (auth *Auth) GetJwtKey() string {
	return config.GetParam("JWT-KEY", "6046ce088ad7283fc513733974f97cbae2f71282").String()
}

// GetJwtTimeOut
// 超时
func (auth *Auth) GetJwtTimeOut() int64 {
	timeOut := config.GetParam("JWT-TIME-OUT", "24").Int64()
	return timeOut
}

// GetJwtMaxRefresh
// 刷新时间
func (auth *Auth) GetJwtMaxRefresh() int64 {
	refresh := config.GetParam("JWT-MAX-REFRESH", "").Int64()
	if refresh == 0 {
		return 5
	} else {
		return refresh
	}
}
