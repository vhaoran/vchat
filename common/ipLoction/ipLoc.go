package ipLoction

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

type (
	CurrentIP struct {
		IP    string `json:"ip"`
		Place string `json:"place"`
	}
)

func NewCurrentIP(ip string) (b *CurrentIP, err error) {
	b = &CurrentIP{IP: ip}
	_ = b.get()
	return
}

func (c *CurrentIP) get() (err error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf("http://ip.taobao.com/service/getIpInfo.php?ip=%s", c.IP)
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("request failed")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if !gjson.ValidBytes(b) {
		return errors.New("request failed,not json")
	}
	rst := gjson.ParseBytes(b)
	code := rst.Get("code").Int()
	if code != 0 {
		return fmt.Errorf("err code = %d", code)
	}
	c.Place = fmt.Sprintf("%s%s%s",
		rst.Get("data.region").String(),
		rst.Get("data.city").String(),
		rst.Get("data.isp").String())
	return
}

// ExtranetIP get external IP addr.
func ExtranetIP() (ip string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("Get external IP error: %v ", p)
		} else if err != nil {
			err = errors.New("Get external IP error: " + err.Error())
		}
	}()
	resp, err := http.Get("http://pv.sohu.com/cityjson?ie=utf-8")
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	idx := bytes.Index(b, []byte(`"cip": "`))
	b = b[idx+len(`"cip": "`):]
	idx = bytes.Index(b, []byte(`"`))
	b = b[:idx]
	ip = string(b)
	return
}

// IntranetIP get internal IP addr.
func IntranetIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("do you connected to the network? ")
}
