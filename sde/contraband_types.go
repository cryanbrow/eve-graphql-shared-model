package sde

type ContrabandType map[int]struct {
	Factions map[int]struct {
		AttackMinSec     float32 `yaml:"attackMinSec" json:"attackMinSec"`
		ConfiscateMinSec float32 `yaml:"confiscateMinSec" json:"confiscateMinSec"`
		FineByValue      float32 `yaml:"fineByValue" json:"fineByValue"`
		StandingLoss     float32 `yaml:"standingLoss" json:"standingLoss"`
	}
	ID int `json:"id"`
}
