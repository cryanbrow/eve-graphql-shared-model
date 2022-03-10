package sde

type Faction map[int]struct {
	CorporationID int `yaml:"corporationID" json:"corporationID"`
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	IconID               int   `yaml:"iconID" json:"iconID"`
	MemberRaces          []int `yaml:"memberRaces" json:"memberRaces"`
	MilitiaCorporationID int   `yaml:"militiaCorporationID" json:"militiaCorporationID"`
	NameID               struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	ShortDescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"shortDescriptionID" json:"shortDescriptionID"`
	SizeFactor    float32 `yaml:"sizeFactor" json:"sizeFactor"`
	SolarSystemID int     `yaml:"solarSystemID" json:"solarSystemID"`
	UniqueName    bool    `yaml:"uniqueName" json:"uniqueName"`
	ID            int     `json:"id"`
}
