package sde

type CategoryID map[int]struct {
	Name struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"name" json:"name"`
	Published bool `yaml:"published" json:"published"`
	ID        int  `json:"id"`
}
