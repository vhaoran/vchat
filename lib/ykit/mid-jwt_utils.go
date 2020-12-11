package ykit

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	tran "github.com/go-kit/kit/transport/http"

	"github.com/vhaoran/vchat/lib/yjwt"
	"github.com/vhaoran/vchat/lib/ylog"
)

const (
	JWT_TOKEN = "Jwt"
	UID_KEY   = "Uid"
)

func Jwt2ctx() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		jwt := req.Header.Get(JWT_TOKEN)
		c := ctx
		if len(jwt) > 0 {
			c = context.WithValue(c, JWT_TOKEN, jwt)
		}

		//

		//
		return c
	}
}

func Jwt2Req() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		c := ctx
		a := ctx.Value(JWT_TOKEN)
		ylog.Debug("----------", "jwt raw", "------------")

		l, ok := a.(string)
		if ok && len(l) > 0 {
			req.Header.Set(JWT_TOKEN, l)
		}

		//

		return c
	}
}

func UIDOfTest(s string) int64 {
	l := []string{"/", "|"}
	for _, v := range l {
		uid := UIDOfTestUnit(s, v)
		if uid > 0 {
			return uid
		}
	}

	return 0
}

func UIDOfTestUnit(s, sign string) int64 {
	if strings.Contains(s, sign) {
		l := strings.Split(s, sign)
		if len(l) > 1 {
			uid, err := strconv.ParseInt(l[1], 10, 64)
			if err != nil {
				ylog.Error("mid-jwt_utils.go->", err)
				return 0
			}

			return uid
		}
		return 0
	}
	return 0
}

func GetUIDOfJwtString(jwt string) int64 {
	if len(jwt) == 0 {
		return 0
	}

	s := jwt
	if len(s) > 0 {
		if i := UIDOfTest(s); i > 0 {
			return i
		}

		uid, err := yjwt.Parse(s)
		if err != nil {
			ylog.Error("mid-jwt_utils.go->", err)
			return 0
		}
		return uid
	}
	return 0
}

func GetUIDOfReq(req *http.Request) int64 {
	s := req.Header.Get(JWT_TOKEN)
	return GetUIDOfJwtString(s)
}

func GetUIDOfContext(ctx context.Context) int64 {
	i := ctx.Value("Uid")

	if i != nil {
		uid, err := strconv.ParseInt(fmt.Sprint(i), 10, 64)
		if err != nil {
			ylog.Error("mid-jwt_utils.go->", err)
			return 0
		}
		return uid
	}

	//--------not found uid then parse jwt -----------------------------
	i = ctx.Value(JWT_TOKEN)
	s, ok := i.(string)
	if !ok {
		return 0
	}

	if len(s) == 0 {
		return 0
	}
	if i := UIDOfTest(s); i > 0 {
		return i
	}

	ii, err := yjwt.Parse(s)
	if err != nil {
		ylog.Error("mid-jwt_utils.go->", err)
		return 0
	}
	return ii
}
