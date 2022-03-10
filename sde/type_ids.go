package sde

type TypeID map[int]struct {
	BasePrice   float64 `yaml:"de" json:"de"`
	Capacity    float32 `yaml:"capacity" json:"capacity"`
	Description struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"description" json:"description"`
	GraphicID     int           `yaml:"GraphicID" json:"GraphicID"`
	GroupID       int           `yaml:"GroupID" json:"GroupID"`
	IconID        int           `yaml:"IconID" json:"IconID"`
	MarketGroupID int           `yaml:"MarketGroupID" json:"MarketGroupID"`
	Masteries     map[int][]int `yaml:"Masteries" json:"Masteries"`
	MetaGroupID   int           `yaml:"MetaGroupID" json:"MetaGroupID"`
	Mass          string        `yaml:"Mass" json:"Mass"`
	Name          struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"name" json:"name"`
	PortionSize    int     `yaml:"PortionSize" json:"PortionSize"`
	Published      bool    `yaml:"Published" json:"Published"`
	RaceID         int     `yaml:"RaceID" json:"RaceID"`
	Radius         float32 `yaml:"Radius" json:"Radius"`
	SofFactionName string  `yaml:"SofFactionName" json:"SofFactionName"`
	SoundID        int     `yaml:"SoundID" json:"SoundID"`
	Traits         Trait   `yaml:"Traits" json:"Traits"`
	Volume         float64 `yaml:"Volume" json:"Volume"`
	ID             int     `json:"id"`
}

type Trait struct {
	IconID      int                 `yaml:"IconID" json:"IconID"`
	MiscBonuses []Bonus             `yaml:"MiscBonuses" json:"MiscBonuses"`
	RoleBonuses []Bonus             `yaml:"RoleBonuses" json:"RoleBonuses"`
	TraitTypes  map[int][]TraitType `yaml:"TraitTypes" json:"TraitTypes"`
}

type Bonus struct {
	BonusAmmount float32 `yaml:"bonus" json:"bonus"`
	BonusText    struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"bonusText" json:"bonusText"`
	Importance int  `yaml:"Importance" json:"Importance"`
	IsPositive bool `yaml:"IsPositive" json:"IsPositive"`
	UnitID     int  `yaml:"UnitID" json:"UnitID"`
}

type TraitType struct {
	BonusAmount float32 `yaml:"bonus" json:"bonus"`
	BonusText   struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"bonusText" json:"bonusText"`
	Importance int `yaml:"Importance" json:"Importance"`
	UnitID     int `yaml:"UnitID" json:"UnitID"`
}
