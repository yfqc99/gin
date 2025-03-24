package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 要编码的字符串
	data := "hello world"

	// 使用标准 Base64 编码
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	// 解码
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("解码错误:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))
}
