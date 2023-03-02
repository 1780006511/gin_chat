package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// EncodeMd5 该方法对上传后的文件名进行格式化，就是对文件名用MD5加密后再进行写入，避免直接暴露原始名称。
func EncodeMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// EncodeMD5 转大写
func EncodeMD5(value string) string {
	return strings.ToUpper(EncodeMd5(value))
}

// MakePassword 加密
func MakePassword(plainpwd, salt string) string {
	return EncodeMd5(plainpwd + salt)
}

// ValidPassword 解密
func ValidPassword(plainpwd, salt string, password string) bool {
	//fmt.Println("EncodeMd5(plainpwd+salt)=", EncodeMd5(plainpwd+salt))
	//fmt.Println("password:", password)
	return EncodeMd5(plainpwd+salt) == password
}
