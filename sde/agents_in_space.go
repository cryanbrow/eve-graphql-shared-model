package sde

type AgentInSpace map[int]struct {
	DungeonID     int `yaml:"dungeonID" json:"dungeonID" xml:"dungeonID"`
	SolarSystemID int `yaml:"solarSystemID" json:"solarSystemID" xml:"solarSystemID"`
	SpawnPointID  int `yaml:"spawnPointID" json:"spawnPointID" xml:"spawnPointID"`
	TypeID        int `yaml:"typeID" json:"typeID" xml:"typeID"`
	ID            int `yaml:"id" json:"id" xml:"id"`
}
