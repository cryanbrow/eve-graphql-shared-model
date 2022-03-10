package sde

type Skin map[int]struct {
	AllowCCPDevs       bool   `yaml:"allowCCPDevs" json:"allowCCPDevs"`
	InternalName       string `yaml:"internalName" json:"internalName"`
	IsStructureSkin    bool   `yaml:"isStructureSkin" json:"isStructureSkin"`
	SkinID             int    `yaml:"skinID" json:"skinID"`
	SkinMaterialID     int    `yaml:"skinMaterialID" json:"skinMaterialID"`
	Types              []int  `yaml:"types" json:"types"`
	VisibleSerenity    bool   `yaml:"visibleSerenity" json:"visibleSerenity"`
	VisibleTranquility bool   `yaml:"visibleTranquility" json:"visibleTranquility"`
	ID                 int    `json:"id"`
}
