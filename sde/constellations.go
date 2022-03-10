package sde

type Constellation struct {
	Center          []float64 `yaml:"center" json:"-"`
	ConstellationID int       `yaml:"constellationID" json:"constellation_id"`
	Systems         []int     `yaml:"-" json:"systems"`
	MaxArray        []float64 `yaml:"max" json:"-"`
	Max             position  `yaml:"-" json:"max"`
	MinArray        []float64 `yaml:"min" json:"-"`
	Min             position  `yaml:"-" json:"min"`
	Name            string    `yaml:"-" json:"name"`
	NameID          int       `yaml:"nameID" json:"name_id"`
	RegionID        int       `yaml:"-" json:"region_id"`
	Position        position  `yaml:"-" json:"position"`
}
