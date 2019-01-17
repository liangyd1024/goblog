//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/18

package crypt

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"testing"
)

var uid, value string

func TestGenerateCaptcha(t *testing.T) {
	uid, value = GenerateCaptcha()
	fmt.Printf("call TestGenerateCaptcha uid:%v,value:%v", uid, value)
}

func TestVerifyCaptcha(t *testing.T) {
	fmt.Printf("call TestVerifyCaptcha uid:%v,value:%v", uid, value)

	image, _ := base64.StdEncoding.DecodeString(strings.Split(value, ",")[1])
	fmt.Println("length:", len(image))
	file, _ := os.Create("/Users/marco/Downloads/verify.jpg")
	defer file.Close()
	bufWriter := bufio.NewWriter(file)

	bufWriter.Write(image)
	bufWriter.Flush()

	fmt.Printf("call TestGenerateCaptcha result:%v", VerifyCaptcha(uid, value))
}
