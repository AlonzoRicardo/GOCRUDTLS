package keygen

import (
	"fmt"
	// "image/png"
	// "os"

	// "github.com/qpliu/qrencode-go/qrencode"
	uuid_gen "github.com/satori/go.uuid"
)

// GenerateKey comments
func GenerateKey() string {
	uuid := uuid_gen.NewV4()
	fmt.Printf("UUIDv4: %s\n", uuid)

	// grid, err := qrencode.Encode(s, qrencode.ECLevelM)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// f, err := os.Create("/Users/ricardoalonzo/go/src/go_auth/qr2.png")

	// if err != nil {
	// 	fmt.Println("error 2")
	// 	return
	// }

	// defer f.Close()
	// png.Encode(f, grid.Image(8))
	return uuid.String()
}
