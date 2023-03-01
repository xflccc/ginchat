package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 小写
func Md5Encode(value string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	tmpStr := hash.Sum(nil)
	return hex.EncodeToString(tmpStr)
}

// MD5Encode  大写
func MD5Encode(value string) string {
	return strings.ToUpper(Md5Encode(value))
}

// MakePassword 加密
func MakePassword(password string, salt string) string {
	return MD5Encode(password + salt)
}

// ValidPassword 解密
func ValidPassword(password string, salt string, oldPassword string) bool {

	return MD5Encode(password+salt) == oldPassword
}

// MakeSalt 生成默认盐值
func MakeSalt() string {
	/*currentTime := time.Now()
	salt := currentTime.Format("2006020130405")*/
	return "ssss"
}
