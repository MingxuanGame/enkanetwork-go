package main

import (
	"fmt"
	"os"

	"github.com/MingxuanGame/enkanetwork-go/api"
	"github.com/MingxuanGame/enkanetwork-go/data/lang"
)

func main() {
	data, err := api.GetInfo(821819570)
	if err != nil {
		fmt.Printf("There is an error when getting data: %v", err)
		os.Exit(1)
	}
	fmt.Printf("---Player Data---\nName: %s\nLevel: %d\nWorld Level: %d\nSignature: %s\nAbyss: %d - %d\n",
		data.PlayerInfo.Nickname,
		data.PlayerInfo.Level,
		data.PlayerInfo.WorldLevel,
		data.PlayerInfo.Signature,
		data.PlayerInfo.TowerFloorIndex,
		data.PlayerInfo.TowerLevelIndex,
	)

	fmt.Println("---Characters---")

	for _, avatar := range data.PlayerInfo.ShowAvatarInfoList {
		nameEN, _ := lang.GetCharacterName(lang.EN, avatar.AvatarID)
		nameCHS, _ := lang.GetCharacterName(lang.CHS, avatar.AvatarID)

		fmt.Printf("Name: %s(Chinese: %s)\nLevel: %d\n------\n",
			nameEN,
			nameCHS,
			avatar.Level,
		)
	}
}
