package sde

type Region struct {
	Center          []float64 `yaml:"center" json:"-"`
	Constellations  []int     `yaml:"-" json:"constellations"`
	DescriptionID   int       `yaml:"descriptionID" json:"description_id"`
	MaxArray        []float64 `yaml:"max" json:"-"`
	Max             position  `yaml:"-" json:"max"`
	MinArray        []float64 `yaml:"min" json:"-"`
	Min             position  `yaml:"-" json:"min"`
	Name            string    `yaml:"-" json:"name"`
	NameID          int       `yaml:"nameID" json:"name_id"`
	Nebula          int       `yaml:"nebula" json:"nebula"`
	RegionID        int       `yaml:"regionID" json:"region_id"`
	Position        position  `yaml:"-" json:"position"`
	WormholeClassID int       `yaml:"wormholeClassID" json:"wormhole_class_id"`
}
