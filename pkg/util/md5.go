package util

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5加密
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}