package model

type PlayerInfo struct {
	Nickname             string `json:"nickname"`
	Level                int    `json:"level"`
	Signature            string `json:"signature"`
	WorldLevel           int    `json:"worldLevel"`
	NameCardID           int    `json:"nameCardId"`
	FinishAchievementNum int    `json:"finishAchievementNum"`
	TowerFloorIndex      int    `json:"towerFloorIndex"`
	TowerLevelIndex      int    `json:"towerLevelIndex"`
	ShowAvatarInfoList   []struct {
		AvatarID  int `json:"avatarId"`
		Level     int `json:"level"`
		CostumeID int `json:"costumeId,omitempty"`
	} `json:"showAvatarInfoList"`
	ShowNameCardIDList []int `json:"showNameCardIdList"`
	ProfilePicture     struct {
		AvatarID int `json:"avatarId"`
	} `json:"profilePicture"`
}
