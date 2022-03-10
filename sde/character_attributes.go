package sde

type CharacterAttribute map[int]struct {
	Description string `yaml:"description" json:"description"`
	IconID      int    `yaml:"iconID" json:"iconID"`
	NameID      struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	Notes            string `yaml:"notes" json:"notes"`
	ShortDescription string `yaml:"shortDescription" json:"shortdescription"`
	ID               int    `json:"id"`
}
