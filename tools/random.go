package tools

import (
	"crypto/md5"
	"fmt"
)

func HashTool(s string) string {
	str := md5.New()
	str.Write([]byte(s))
	res := fmt.Sprintf("%x", str.Sum(nil))
	return res
}

func GenerateRandom() {

}
