package sde

type Ancestry map[int]struct {
	BloodlineID   int `yaml:"bloodlineID" json:"bloodlineID" xml:"bloodlineID"`
	Charisma      int `yaml:"charisma" json:"charisma" xml:"charisma"`
	DescriptionID struct {
		DE string `yaml:"de" json:"de" xml:"de"`
		EN string `yaml:"en" json:"en" xml:"en"`
		FR string `yaml:"fr" json:"fr" xml:"fr"`
		JA string `yaml:"ja" json:"ja" xml:"ja"`
		KO string `yaml:"ko" json:"ko" xml:"ko"`
		RU string `yaml:"ru" json:"ru" xml:"ru"`
		ZH string `yaml:"zh" json:"zh" xml:"zh"`
	} `yaml:"descriptionID" json:"descriptionID" xml:"descriptionID"`
	IconID       int `yaml:"iconID" json:"iconID" xml:"iconID"`
	Intelligence int `yaml:"intelligence" json:"intelligence" xml:"intelligence"`
	Memory       int `yaml:"memory" json:"memory" xml:"memory"`
	NameID       struct {
		DE string `yaml:"de" json:"de" xml:"de"`
		EN string `yaml:"en" json:"en" xml:"en"`
		FR string `yaml:"fr" json:"fr" xml:"fr"`
		JA string `yaml:"ja" json:"ja" xml:"ja"`
		KO string `yaml:"ko" json:"ko" xml:"ko"`
		RU string `yaml:"ru" json:"ru" xml:"ru"`
		ZH string `yaml:"zh" json:"zh" xml:"zh"`
	} `yaml:"nameID" json:"nameID" xml:"nameID"`
	Perception       int    `yaml:"perception" json:"perception" xml:"perception"`
	ShortDescription string `yaml:"shortDescription" json:"shortDescription" xml:"shortDescription"`
	Willpower        int    `yaml:"willpower" json:"willpower" xml:"willpower"`
	ID               int    `yaml:"id" json:"id" xml:"id"`
}
