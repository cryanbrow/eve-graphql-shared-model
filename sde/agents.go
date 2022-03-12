package sde

type Agent map[int]struct {
	AgentTypeID   int  `yaml:"agentTypeID" json:"agentTypeID" xml:"agentTypeID"`
	CorporationID int  `yaml:"corporationID" json:"corporationID" xml:"corporationID"`
	DivisionID    int  `yaml:"divisionID" json:"divisionID" xml:"divisionID"`
	IsLocator     bool `yaml:"isLocator" json:"isLocator" xml:"isLocator"`
	Level         int  `yaml:"level" json:"level" xml:"level"`
	LocationID    int  `yaml:"locationID" json:"locationID" xml:"locationID"`
	ID            int  `yaml:"id" json:"id" xml:"id"`
}
