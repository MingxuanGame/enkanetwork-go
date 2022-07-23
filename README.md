# EnkaNetwork-Go

**EN** | [CHS](README_CHS.md)

🎉This project is a wrapper for https://enka.network written in Go.

Godoc: https://pkg.go.dev/github.com/MingxuanGame/enkanetwork-go

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

Output:

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

You can get this on [example/overview.go](example/overview.go).

## 💡Example

See the [example](example/) folder.

## 🌐Localization(l10n)

You can get characters' name via the function `GetCharacterName` from [data/lang](data/lang/) package.

For example, see this program, it will print a character's all name.(You can find it on [example/l10n.go](example/l10n.go))

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

Output:

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

The output language order is related to the following table.

| Language Const | Language Name |
|----------------|---------------|
| EN             | English       |
| RU             | русский       |
| VI             | Tiếng Việt    |
| TH             | ไทย           |
| PT             | português     |
| KR             | 한국어        |
| JP             | 日本語        |
| ID             | Indonesian    |
| FR             | français      |
| ES             | español       |
| DE             | deutsch       |
| CHT            | 正體中文      |
| CHS            | 简体中文      |

For the `Language Const` in the above table, you can directly use `lang.[Language Const]` to get it, such as `lang.CHS`.

## 📄LICENSE

[MIT License](LICENSE)
