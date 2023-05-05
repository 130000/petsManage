package main

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secretKey")

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusOK, CommonResult{unauthorized, "权限不足", nil})
			c.Abort()
			return
		}
		token = token[7:]
		claims, OK := parseToken(token)
		if !OK {
			c.JSON(http.StatusOK, CommonResult{unauthorized, "权限不足", nil})
			c.Abort()
			return
		}
		userId := claims.UserId
		err1, _ := findUserById(userId)
		if err1 == sql.ErrNoRows {
			c.JSON(http.StatusOK, CommonResult{unauthorized, "权限不足", nil})
			c.Abort()
			return
		} else if err1 != nil {
			c.JSON(http.StatusOK, errHandle(err1))
			c.Abort()
			return
		}
		c.Next()
	}
}
func generateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			Issuer:    "test"}})
	tokenString, err := token.SignedString(jwtKey)
	return "Bearer " + tokenString, err
}
func parseToken(token string) (*MyCustomClaims, bool) {
	claims := &MyCustomClaims{}
	newToken, _ := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if key, OK := newToken.Claims.(*MyCustomClaims); OK && newToken.Valid {
		return key, true
	}
	return nil, false
}
