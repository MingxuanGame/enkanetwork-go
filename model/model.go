// This package defines the model returned by Enka.Network
// Please also use with Enka.Network documentation  
// 
// https://api.enka.network
package model


// playerInfo JSON-data
// https://api.enka.network/#/api?id=playerinfo
type PlayerInfo struct {
	Nickname             string `json:"nickname"`
	Level                uint8  `json:"level"`
	Signature            string `json:"signature"`
	WorldLevel           uint8  `json:"worldLevel"`
	NameCardID           uint32 `json:"nameCardId"`
	FinishAchievementNum uint32 `json:"finishAchievementNum"`
	TowerFloorIndex      uint8  `json:"towerFloorIndex"` // max: 12
	TowerLevelIndex      uint8  `json:"towerLevelIndex"` // max: 3
	ShowAvatarInfoList   []struct {
		AvatarID  uint32 `json:"avatarId"`
		Level     uint8  `json:"level"`
		CostumeID uint32 `json:"costumeId,omitempty"`
	} `json:"showAvatarInfoList"`
	ShowNameCardIDList []uint32 `json:"showNameCardIdList,omitempty"`
	ProfilePicture     struct {
		AvatarID uint32 `json:"avatarId"`
	} `json:"profilePicture"`
}

// ===== Characters Info ===== //

// propMap JSON-data
// https://api.enka.network/#/api?id=propmap
type PropMap struct {
	Type uint32 `json:"type"`
	Ival string `json:"ival"` // Ignore It!
	Val  string `json:"val,omitempty"`
}

// reliquary JSON-data
// https://api.enka.network/#/api?id=reliquary
type Artifact struct {
	Level            uint8    `json:"level"`
	MainPropId       uint16   `json:"mainPropId"`
	AppendPropIdList []uint32 `json:"appendPropIdList"`
}

// weapon JSON-data
// https://api.enka.network/#/api?id=weapon
type Weapon struct {
	Level        uint8            `json:"level"`
	PromoteLevel uint8            `json:"promoteLevel"`
	AffixMap     map[string]uint8 `json:"affixMap"`
}

// reliquaryMainstat, reliquarySubstats, weaponStats JSON-data
// https://api.enka.network/#/api?id=reliquarymainstat-reliquarysubstats-weaponstats
//
// SubPropId: appendPropId
//
// Value: propValue
type Stat struct {
	MainPropId string  `json:"mainPropId,omitempty"`
	SubPropId  string  `json:"appendPropId,omitempty"`
	Value      float64 `json:"statValue"`
}

// flat JSON-data
// https://api.enka.network/#/api?id=flat
// 
// NameTextHash: nameTextHashMap
//
// SetNameTextHash: setNameTextHashMap
type Flat struct {
	// l10n
	NameTextHash    uint32 `json:"nameTextHashMap"`
	SetNameTextHash uint32 `json:"setNameTextHashMap,omitempty"`

	// artifact
	ReliquaryMainStat Stat   `json:"reliquaryMainstat,omitempty"`
	ReliquarySubStats []Stat `json:"reliquarySubstats,omitempty"`
	EquipType         string `json:"equipType,omitempty"`

	// weapon
	WeaponStat []Stat `json:"weaponStats,omitempty"`

	RankLevel uint8  `json:"rankLevel"` // 3, 4 or 5
	ItemType  string `json:"itemType"`  // ITEM_WEAPON or ITEM_RELIQUARY
	Icon      string `json:"icon"`      // You can get the icon from https://enka.network/ui/{Icon}.png
}

// every equipList element
// https://api.enka.network/#/api?id=equiplist
type Equip struct {
	ItemId    uint32   `json:"itemId"`
	Reliquary Artifact `json:"reliquary,omitempty"`
	Weapon    Weapon   `json:"weapon"`
	Flat      Flat     `json:"flat"`
}

// every avatarInfoList element
// https://api.enka.network/#/api?id=avatarinfolist
type AvatarInfo struct {
	AvatarId               uint32             `json:"avatarID"`
	TalentIdList           []uint16           `json:"talentIdList,omitempty"`
	PropMap                map[string]PropMap `json:"propMap"`
	FightPropMap           map[string]float64 `json:"fightPropMap"`
	SkillDepotId           uint16             `json:"skillDepotId"`
	InherentProudSkillList []uint32           `json:"inherentProudSkillOpens"`
	SkillLevelMap          map[string]uint32  `json:"skillLevelMap"`
	EquipList              []Equip            `json:"equipList"`
	FetterInfo             struct {
		ExpLevel uint8 `json:"expLevel"`
	} `json:"fetterInfo"`
}

// avatarInfoList JSON-data
// https://api.enka.network/#/api?id=avatarinfolist
type AvatarInfoList []AvatarInfo

// all JSON-data
// https://api.enka.network/#/api?id=data-structure-info
type EnkaNetworkData struct {
	PlayerInfo     PlayerInfo     `json:"playerInfo"`
	AvatarInfoList AvatarInfoList `json:"avatarInfoList"`
}
