package main

import (
	"fmt"
	"image/color"

	"github.com/skip2/go-qrcode"
)

func main() {
	// err := qrcode.WriteFile("hello this is qrcode in golang", qrcode.Medium, 256, "myfirst_file.png")
	// if err != nil {
	// 	fmt.Printf("Sorry couldn't create qrcode:,%v", err)
	// }

	err := qrcode.WriteColorFile("this is colored", qrcode.Medium, 256, color.Black, color.White, "secondfile.png")
	if err != nil {
		fmt.Printf("Sorry couldn't create qrcode:,%v", err)

	}
}
