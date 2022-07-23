# EnkaNetwork-Go

[EN](README.md) | **CNS**

🎉这个项目是使用 Go 编写的 https://enka.network API 包装

## 安装

```bash
go get github.com/MingxuanGame/enkanetwork-go
```

## 🚀使用方法

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

输出：

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

## 💡示例

请查看 [example](example/) 文件夹

## 📄许可证

[MIT License](LICENSE)
