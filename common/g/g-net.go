package g

import (
	"net/http"
	"strings"
)

//根据req，获取客户端请求的ip
func GetClientIP(req *http.Request) string {
	s := req.Header.Get("X-Real-IP")
	if len(s) == 0 {
		s = req.Header.Get("X-Forwarded-For")
	}

	return cutIP(s)
}

func cutIP(ip string) string {
	s := ip
	if len(s) == 0 {
		return ""
	}

	i := strings.Index(s, ":")
	if i >= 0 && i < len(s) {
		l := strings.Split(s, ":")
		if len(l) > 0 {
			s = l[0]
		}
	}

	return s
}
