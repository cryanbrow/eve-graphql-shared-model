package sde

type ControlTowerAttribute map[int]struct {
	Resources []Resource `yaml:"resource" json:"resource"`
	ID        int        `json:"id"`
}

type Resource struct {
	FactionID        int     `yaml:"factionID" json:"factionID"`
	MinSecurityLevel float32 `yaml:"minSecurityLevel" json:"minSecurityLevel"`
	Purpose          int     `yaml:"purpose" json:"purpose"`
	Quantity         int     `yaml:"quantity" json:"quantity"`
	ResourceTypeID   int     `yaml:"resourceTypeID" json:"resourceTypeID"`
}
