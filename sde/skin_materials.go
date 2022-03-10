package sde

type SkinMaterial map[int]struct {
	DisplayNameID  int `yaml:"displayNameID" json:"displayNameID"`
	M3aterialSetID int `yaml:"materialSetID" json:"materialSetID"`
	SkinMaterialID int `yaml:"skinMaterialID" json:"skinMaterialID"`
	ID             int `json:"id"`
}
