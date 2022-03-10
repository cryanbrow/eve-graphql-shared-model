package sde

type ResearchAgent map[int]struct {
	Skills []struct {
		TypeID int `yaml:"typeID" json:"typeID"`
	} `yaml:"skils" json:"skills"`
	ID int `json:"id"`
}
