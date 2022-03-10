package sde

type Landmark map[int]struct {
	DescriptionID  int       `yaml:"descriptionID" json:"descriptionID"`
	IconID         int       `yaml:"iconID" json:"iconID"`
	LandmarkNameID int       `yaml:"landmarkNameID" json:"landmarkNameID"`
	LocationID     int       `yaml:"locationID" json:"locationID"`
	Position       []float64 `yaml:"position" json:"position"`
	ID             int       `json:"id"`
}
