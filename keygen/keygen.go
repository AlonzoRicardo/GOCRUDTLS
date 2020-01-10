package keygen

import (
	"fmt"
	"image/png"
	"os"

	"github.com/qpliu/qrencode-go/qrencode"
)

func GenerateKey(s string) {
	grid, err := qrencode.Encode(s, qrencode.ECLevelM)
	fmt.Println(grid)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Create("/Users/ricardoalonzo/go/src/go_auth/qr2.png")
	fmt.Println(f)
	if err != nil {
		fmt.Println("error 2")
		return
	}
	defer f.Close()
	png.Encode(f, grid.Image(8))
	return
}
