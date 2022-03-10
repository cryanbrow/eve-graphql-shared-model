package sde

type DogmaAttributeCategory map[int]struct {
	Description string `yaml:"description" json:"description"`
	Name        string `yaml:"name" json:"name"`
	ID          int    `json:"id"`
}
