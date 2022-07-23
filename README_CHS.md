# EnkaNetwork-Go

[EN](README.md) | **CNS**

🎉这个项目是使用 Go 编写的 https://enka.network API 包装

Godoc: https://pkg.go.dev/github.com/MingxuanGame/enkanetwork-go
## 🚩安装

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
Name: Keqing(Chinese: 刻晴)
Level: 60
------
Name: Mona(Chinese: 莫娜)
Level: 60
------
Name: Noelle(Chinese: 诺艾尔)
Level: 35
------
Name: Diona(Chinese: 迪奥娜)
Level: 58
------
Name: Traveler(Chinese: 旅行者)
Level: 50
------
```

你可以在 [example/overview.go](example/overview.go) 获取此程序

## 💡示例

请查看 [example](example/) 文件夹

## 🌐本地化（l10n）

你可以通过包 [data/lang](data/lang/) 中的函数 `GetCharacterName` 获取角色名称

例如下面的程序，它将会打印一个角色所有语言的名称（你可以在 [example/l10n.go](example/l10n.go) 找到它）

```go
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
```

输出：

```
Keqing
Кэ Цин
Keqing
Keqing
Keqing
각청
刻晴
Keqing
Keqing
Keching
Keqing
刻晴
刻晴
```

输出语言的顺序与下表相关

| 语言常量 | 语言名称   |
|----------|------------|
| EN       | English    |
| RU       | русский    |
| VI       | Tiếng Việt |
| TH       | ไทย        |
| PT       | português  |
| KR       | 한국어     |
| JP       | 日本語     |
| ID       | Indonesian |
| FR       | français   |
| ES       | español    |
| DE       | deutsch    |
| CHT      | 正體中文   |
| CHS      | 简体中文   |

对于上表的`语言常量`，你可以通过 `lang.[语言常量]` 获取它，例如 `lang.CHS`

## 📄许可证

[MIT License](LICENSE)
