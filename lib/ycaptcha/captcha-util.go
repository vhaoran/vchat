package ycaptcha

import (
	cap "github.com/dchest/captcha"
	"github.com/vhaoran/vchat/lib/ylog"
	"net/http"
)

func Init() {
	cap.SetCustomStore(new(MyStore))
	ylog.Debug(".......ycaptcha init........")
}

//
func GetCaptchatID() string {
	return cap.New()
}

func Verify(id, code string) bool {
	return cap.VerifyString(id, code)
}

func Handler() http.Handler {
	return cap.Server(cap.StdWidth, cap.StdHeight)
}
