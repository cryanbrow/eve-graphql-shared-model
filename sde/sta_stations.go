package sde

type StaStation struct {
	ConstellationID          int     `yaml:"constellationID" json:"constellation_id"`
	CorporationID            int     `yaml:"corporationID" json:"corporation_id"`
	DockingCostPerVolume     int     `yaml:"dockingCostPerVolume" json:"docking_cost_per_volume"`
	MaxShipVolumeDockable    int     `yaml:"maxShipVolumeDockable" json:"max_ship_volume_dockable"`
	OfficeRentalCost         int     `yaml:"officeRentalCost" json:"office_rental_cost"`
	OperationID              int     `yaml:"operationID" json:"operation_id"`
	RegionID                 int     `yaml:"regionID" json:"region_id"`
	ReprocessingEfficiency   float32 `yaml:"reprocessingEfficiency" json:"reprocessing_efficiency"`
	ReprocessingHangarFlag   int     `yaml:"reprocessingHangarFlag" json:"reprocessing_hangar_flag"`
	ReprocessingStationsTake float32 `yaml:"reprocessingStationsTake" json:"reprocessing_stations_take"`
	Security                 float32 `yaml:"security" json:"security"`
	SolarSystemID            int     `yaml:"solarSystemID" json:"solar_system_id"`
	StationID                int     `yaml:"stationID" json:"station_id"`
	StationName              string  `yaml:"stationName" json:"station_name"`
	StationTypeID            int     `yaml:"stationTypeID" json:"station_type_id"`
	X                        float64 `yaml:"x" json:"x"`
	Y                        float64 `yaml:"y" json:"y"`
	Z                        float64 `yaml:"z" json:"z"`
}
