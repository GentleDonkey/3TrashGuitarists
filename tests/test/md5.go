package main

import (
	"3TrashGuitarists/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMd5(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))

	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	md5_pwd := GetMd5("19971004a")
	fmt.Println(md5_pwd)

	fmt.Println(utils.VerifyNameFormat("serene"))
	fmt.Println(utils.VerifyPasswordFormat("19971004a"))
}
