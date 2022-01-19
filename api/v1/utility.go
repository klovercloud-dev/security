package v1

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

func getUserResourcePermissionFromBearerToken(context echo.Context, u userApi) (v1.UserResourcePermission, error) {
	bearerToken := context.Request().Header.Get("Authorization")
	if bearerToken == "" {
		return  v1.UserResourcePermission{}, errors.New("[ERROR]: No token found!")
	}
	var token string
	if len(strings.Split(bearerToken, " ")) == 2 {
		token = strings.Split(bearerToken, " ")[1]
	} else {
		return  v1.UserResourcePermission{}, errors.New( "[ERROR]: No token found!")
	}
	if !u.jwtService.IsTokenValid(token) {
		return v1.UserResourcePermission{},errors.New("[ERROR]: Token is expired!")
	}
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Publickey), nil
	})
	jsonbody, err := json.Marshal(claims["data"])
	if err != nil {
		log.Println(err)
	}
	userResourcePermission := v1.UserResourcePermission{}
	if err := json.Unmarshal(jsonbody, &userResourcePermission); err != nil {
		return v1.UserResourcePermission{}, errors.New("[ERROR]: No resource permissions!")
	}
	return userResourcePermission, nil
}
