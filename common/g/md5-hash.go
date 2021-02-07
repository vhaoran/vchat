package g

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
)

func MD5(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
func HashByte(s []byte) string {
	if len(s) == 0 {
		return ""
	}

	h := sha1.New()
	h.Write(s)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func FileHash(path string) (string, error) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	hash := HashByte(buffer)
	return hash, nil
}
