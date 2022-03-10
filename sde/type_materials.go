package sde

type TypeMaterial map[int]struct {
	Materials []struct {
		MaterialTypeID int `yaml:"materialTypeID" json:"materialTypeID"`
		Quantity       int `yaml:"quantity" json:"quantity"`
	} `yaml:"materials" json:"materials"`
	ID int `json:"id"`
}
