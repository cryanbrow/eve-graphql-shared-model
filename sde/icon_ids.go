package sde

type IconID map[int]struct {
	Description string `yaml:"description" json:"description"`
	IconFile    string `yaml:"iconFile" json:"iconFile"`
	ID          int    `json:"id"`
}
