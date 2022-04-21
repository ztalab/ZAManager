package jwt

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JwtExpireSecend = 24 * 60 * 60

var Secret = []byte("JwtSecret-ZTA")

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
	//Password string `json:"password"`
	//RoleID     []int  `json:"role_id"` // 角色ID
	//BufferTime int64
	jwt.StandardClaims
}

func GenerateToken(id int64, username string) (string, error) {
	nowTime := time.Now()
	// TODO 设置token默认过期时间为24小时
	expireTime := nowTime.Add(JwtExpireSecend * time.Second)
	claims := Claims{
		ID:       id,
		Username: username,
		//Password: utils.NewMd5(password),
		//RoleID:   nil, // 角色ID
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "MSP",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(Secret)
	return token, err
}

func ParseTokenWithClaims(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func ParseTokenV2(token string) (*jwt.Token, error) {
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	return tokenClaims, err
}

func GetIDFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
