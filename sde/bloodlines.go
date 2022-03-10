package sde

type Bloodline map[int]struct {
	Charisma      int `yaml:"charisma" json:"charisma"`
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
	IconID       int `yaml:"iconID" json:"iconID"`
	Intelligence int `yaml:"intelligence" json:"intelligence"`
	Memory       int `yaml:"memory" json:"memory"`
	NameID       struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	Perception int `yaml:"perception" json:"perception"`
	RaceID     int `yaml:"raceID" json:"raceID"`
	Willpower  int `yaml:"willpower" json:"willpower"`
	ID         int `json:"id"`
}
