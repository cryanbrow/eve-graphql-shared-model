package sde

type SkinLicense map[int]struct {
	Duration      int `yaml:"duration" json:"duration"`
	LicenseTypeID int `yaml:"licenseTypeID" json:"licenseTypeID"`
	SkinID        int `yaml:"skinID" json:"skinID"`
	ID            int `json:"id"`
}
