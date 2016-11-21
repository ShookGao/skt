package skt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
	"runtime"
)

// IsExist  whether a file or directory exists
func IsExist(name string) bool {
	_, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsWindows whether Windows system
func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

// IsBlank whether blank
func IsBlank(value reflect.Value) bool {
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// MD5 crypto
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// ToString 转换任意值为string
func ToString(i interface{}) string {
	return fmt.Sprint(i)
}

// ToByte 转换任意值为[]byte
func ToByte(i interface{}) []byte {
	return []byte(fmt.Sprint(i))
}
