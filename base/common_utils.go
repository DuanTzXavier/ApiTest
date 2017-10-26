package base

import (
	"io"
	"crypto/rand"
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
)

func GenerateUniqueCode(digitsa int) string {
	//TODO 生成digitsa位唯一值
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	h := md5.New()
	h.Write([]byte(base64.URLEncoding.EncodeToString(b)))

	return hex.EncodeToString(h.Sum(nil))
}
