package main

import qrcode "github.com/skip2/go-qrcode"
import "fmt"

func main() {
	err := qrcode.WriteFile("https://shop.zk-mall.net?userId=556", qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("write error")
	}
}
