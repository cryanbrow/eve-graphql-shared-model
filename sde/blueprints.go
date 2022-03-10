package sde

type Blueprint map[int]struct {
	Activites struct {
		Copying struct {
			Time int `yaml:"time" json:"time"`
		} `yaml:"copying" json:"copying"`
		Invention struct {
			Materials []Material `yaml:"materials" json:"materials"`
			Products  []Product  `yaml:"products" json:"products"`
			Skills    []Skill    `yaml:"skills" json:"skills"`
			Time      int        `yaml:"time" json:"time"`
		} `yaml:"invention" json:"invention"`
		Manufacturing struct {
			Materials []Material `yaml:"materials" json:"materials"`
			Products  []Product  `yaml:"products" json:"products"`
			Skills    []Skill    `yaml:"skills" json:"skills"`
			Time      int        `yaml:"time" json:"time"`
		} `yaml:"manufacturing" json:"manufacturing"`
		ResearchMaterial struct {
			Time int `yaml:"time" json:"time"`
		} `yaml:"research_material" json:"research_material"`
		ResearchTime struct {
			Time int `yaml:"time" json:"time"`
		} `yaml:"research_time" json:"research_time"`
	} `yaml:"activities" json:"activities"`
	BlueprintTypeID    int `yaml:"blueprintTypeID" json:"blueprintTypeID"`
	MaxProductionLimit int `yaml:"maxProductionLimit" json:"maxProductionLimit"`
	ID                 int `json:"id"`
}

type Material struct {
	Quantity int `yaml:"quantity" json:"quantity"`
	TypeID   int `yaml:"typeID" json:"typeID"`
}

type Product struct {
	Probability float32 `yaml:"probability" json:"probability"`
	Quantity    int     `yaml:"quantity" json:"quantity"`
	TypeID      int     `yaml:"typeID" json:"typeID"`
}

type Skill struct {
	Level  int `yaml:"level" json:"level"`
	TypeID int `yaml:"typeID" json:"typeID"`
}
