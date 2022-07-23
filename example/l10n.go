package main

import (
	"fmt"
	"os"

	"github.com/MingxuanGame/enkanetwork-go/data/lang"
)

func main() {
	var avatarID uint32 = 10000042 // Keqing

	for n := 0; n < 13; n++ {
		name, err := lang.GetCharacterName(lang.Lang(n), avatarID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(name)
	}
}
