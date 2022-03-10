package sde

type AgentInSpace map[int]struct {
	DungeonID     int `yaml:"dungeonID" json:"dungeonID"`
	SolarSystemID int `yaml:"solarSystemID" json:"solarSystemID"`
	SpawnPointID  int `yaml:"spawnPointID" json:"spawnPointID"`
	TypeID        int `yaml:"typeID" json:"typeID"`
	ID            int `json:"id"`
}
