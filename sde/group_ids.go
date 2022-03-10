package sde

type GroupID map[int]struct {
	Anchorable           bool `yaml:"anchorable" json:"anchorable"`
	Anchored             bool `yaml:"anchored" json:"anchored"`
	CategoryID           int  `yaml:"categoryID" json:"categoryID"`
	FittableNonSingleton bool `yaml:"fittableNonSingleton" json:"fittableNonSingleton"`
	Name                 struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"name" json:"name"`
	Published    bool `yaml:"published" json:"published"`
	UseBasePrice bool `yaml:"useBasePrice" json:"useBasePrice"`
	ID           int  `json:"id"`
}
