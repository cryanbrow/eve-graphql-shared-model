package sde

type PlanetSchematic map[int]struct {
	CycleTime int `yaml:"cycleTime" json:"cycleTime"`
	NameID    struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	Pins  []int                `yaml:"pins" json:"pins"`
	Types map[int]ResourceType `yaml:"types" json:"types"`
	ID    int                  `json:"id"`
}

type ResourceType struct {
	IsInput  bool `yaml:"isInput" json:"isInput"`
	Quantity int  `yaml:"quantity" json:"quantity"`
}
