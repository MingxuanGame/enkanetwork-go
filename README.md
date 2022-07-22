# EnkaNetwork-Go

🎉This project is a wrapper for https://enka.network written in Go.

## 🚩Installation

```bash
go get github.com/MingxuanGame/enkanetwork-go
```

## 🚀Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/MingxuanGame/enkanetwork-go/api"
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
	nameMap := map[int]string{
		10000042: "Keqing",
		10000041: "Mona",
	}

	for _, avatar := range data.PlayerInfo.ShowAvatarInfoList {
		fmt.Printf("Name: %s\nLevel: %d\n------\n", nameMap[int(avatar.AvatarID)], avatar.Level)
	}
}
```

Output:

```
---Player Data---
Name: MingxuanGame
Level: 38
World Level: 4
Signature: Do want you like.
Abyss: 4 - 2
---Characters---
Name: Keqing    
Level: 60       
------
Name: Mona      
Level: 60       
------
```

## 💡Example

See the [example](example/) folder.

## 📄LICENSE

[MIT License](LICENSE)