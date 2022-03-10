package sde

type Agent map[int]struct {
	AgentTypeID   int  `yaml:"agentTypeID" json:"agentTypeID"`
	CorporationID int  `yaml:"corporationID" json:"corporationID"`
	DivisionID    int  `yaml:"divisionID" json:"divisionID"`
	IsLocator     bool `yaml:"isLocator" json:"isLocator"`
	Level         int  `yaml:"level" json:"level"`
	LocationID    int  `yaml:"locationID" json:"locationID"`
	ID            int  `json:"id"`
}
