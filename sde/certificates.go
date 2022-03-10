package sde

type Certificate map[int]struct {
	Description    string `yaml:"description" json:"description"`
	GroupID        string `yaml:"groupID" json:"groupID"`
	Name           string `yaml:"name" json:"name"`
	RecommendedFor []int  `yaml:"recommendedFor" json:"recommendedFor"`
	SkillTypes     struct {
		SkillType map[int]struct {
			Advanced int `yaml:"advanced" json:"advanced"`
			Basic    int `yaml:"basic" json:"basic"`
			Elite    int `yaml:"elite" json:"elite"`
			Improved int `yaml:"improved" json:"improved"`
			Standard int `yaml:"standard" json:"standard"`
		}
	}
	ID int `json:"id"`
}
