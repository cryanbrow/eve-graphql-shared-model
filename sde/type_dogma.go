package sde

type TypeDogma map[int]struct {
	DogmaAttributes []struct {
		AttributeID int     `yaml:"attributeID" json:"attributeID"`
		Value       float32 `yaml:"value" json:"value"`
	} `yaml:"dogmaAttributes" json:"dogmaAttributes"`
	DogmaEffects []struct {
		EffectID  int  `yaml:"effectID" json:"effectID"`
		IsDefault bool `yaml:"isDefault" json:"isDefault"`
	} `yaml:"dogmaEffects" json:"dogmaEffects"`
	ID int `json:"id"`
}
