api/user.go
// 仅有登录部分有改动
func login(c *gin.Context) {
if err := c.ShouldBind(&model.User{}); err != nil {
utils.RespFail(c, "verification failed")
return
}
// 传入用户名和密码
username := c.PostForm("username")
password := c.PostForm("password")

// 验证用户名是否存在
flag := dao.SelectUser(username)
// 不存在则退出
if !flag {
// 以 JSON 格式返回信息
utils.RespFail(c, "user doesn't exists")
return
}

// 查找正确的密码
selectPassword := dao.SelectPasswordFromUsername(username)
// 若不正确则传出错误
if selectPassword != password {
// 以 JSON 格式返回信息
utils.RespFail(c, "wrong password")
return
}

// 正确则登录成功
// 创建一个我们自己的声明
claim := model.MyClaims{
Username: username, // 自定义字段
StandardClaims: jwt.StandardClaims{
ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
Issuer:    "Yxh",                                // 签发人
},
}
// 使用指定的签名方法创建签名对象
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
// 使用指定的secret签名并获得完整的编码后的字符串token
tokenString, _ := token.SignedString(middleware.Secret)
utils.RespSuccess(c, tokenString)
}
api/middleware/jwt.go
package middleware

import (
"errors"
"net/http"
"strings"

"gin-demo/model"
"github.com/dgrijalva/jwt-go"
"github.com/gin-gonic/gin"
)

var Secret = []byte("YXH")

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
return func(c *gin.Context) {
// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
// 这里假设Token放在Header的Authorization中，并使用Bearer开头
// 这里的具体实现方式要依据你的实际业务情况决定
authHeader := c.Request.Header.Get("Authorization")
if authHeader == "" {
c.JSON(http.StatusOK, gin.H{
"code": 2003,
"msg":  "请求头中auth为空",
})
c.Abort()
return
}
// 按空格分割
parts := strings.SplitN(authHeader, " ", 2)
if !(len(parts) == 2 && parts[0] == "Bearer") {
c.JSON(http.StatusOK, gin.H{
"code": 2004,
"msg":  "请求头中auth格式有误",
})
c.Abort()
return
}
// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
mc, err := ParseToken(parts[1])
if err != nil {
c.JSON(http.StatusOK, gin.H{
"code": 2005,
"msg":  "无效的Token",
})
c.Abort()
return
}
// 将当前请求的username信息保存到请求的上下文c上
c.Set("username", mc.Username)
c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*model.MyClaims, error) {
// 解析token
token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
return Secret, nil
})
if err != nil {
return nil, err
}
if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid { // 校验token
return claims, nil
}
return nil, errors.New("invalid token")
}