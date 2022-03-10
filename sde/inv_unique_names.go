package sde

type UniqueNames struct {
	GroupID  int    `yaml:"groupID" json:"groupID"`
	ItemID   int    `yaml:"itemID" json:"itemID"`
	ItemName string `yaml:"itemName" json:"itemName"`
}

type NameForID struct {
	GroupID  int    `yaml:"groupID" json:"groupID"`
	ItemName string `yaml:"itemName" json:"itemName"`
}

type IDForName struct {
	GroupID int `yaml:"groupID" json:"groupID"`
	ItemID  int `yaml:"itemID" json:"itemID"`
}
