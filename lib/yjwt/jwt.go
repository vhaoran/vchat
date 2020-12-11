package yjwt

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/robbert229/jwt"

	"github.com/vhaoran/vchat/lib/yconfig"
)

var (
	jwtSecretKey string = "vchat_jwt_key"
	jwtExpired          = 24 * 365 * 10
)

func InitJwt(cfg yconfig.JwtConfig) error {
	if cfg.Expired > 0 {
		jwtExpired = cfg.Expired
	}
	if len(cfg.SecretKey) > 0 {
		jwtSecretKey = cfg.SecretKey
	}

	return nil
}

func Gen(uid int64) (token string, err error) {
	return jwtGen(jwtSecretKey, uid)
}
func Parse(token string) (uid int64, err error) {
	return jwtParse(jwtSecretKey, token)
}

func jwtGen(secretKey string, uid int64) (string, error) {
	algorithm := jwt.HmacSha256(secretKey)

	claims := jwt.NewClaim()
	claims.Set("uid", fmt.Sprint(uid))
	//claims.SetTime("exp", time.Now().Add(time.Hour*24*365*10))

	token, err := algorithm.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func jwtParse(secretKey, token string) (uid int64, err error) {
	uid, err = 0, nil

	algorithm := jwt.HmacSha256(secretKey)
	//
	if err = algorithm.Validate(token); err != nil {
		return
	}

	//parse
	loadedClaims, err := algorithm.Decode(token)
	if err != nil {
		return 0, err
	}

	id, err := loadedClaims.Get("uid")
	if err != nil {
		return 0, err
	}

	s, ok := id.(string)
	if !ok {
		log.Println("uid:", uid)
		return 0, errors.New("错误的uid类型")
	}
	if uid, err = strconv.ParseInt(s, 10, 64); err != nil {
		log.Println("获取jwt时出错,err:", err)
		return 0, err
	}
	return
}
