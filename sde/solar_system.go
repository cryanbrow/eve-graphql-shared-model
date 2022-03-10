package sde

type SolarSystem struct {
	Border                     bool             `yaml:"border" json:"border"`
	Center                     []float64        `yaml:"center" json:"-"`
	ConstellationID            int              `yaml:"-" json:"constellation_id"`
	Corridor                   bool             `yaml:"corridor" json:"corridor"`
	DisallowedAnchorCategories []int            `yaml:"disallowedAnchorCategories" json:"disallowed_anchor_categories"`
	DisallowedAnchorGroups     []int            `yaml:"disallowedAnchorGroups" json:"disallowed_anchor_groups"`
	Fringe                     bool             `yaml:"fringe" json:"fringe"`
	Hub                        bool             `yaml:"hub" json:"hub"`
	International              bool             `yaml:"international" json:"international"`
	Luminosity                 float64          `yaml:"luminosity" json:"luminosity"`
	MaxArray                   []float64        `yaml:"max" json:"-"`
	Max                        position         `yaml:"-" json:"max"`
	MinArray                   []float64        `yaml:"min" json:"-"`
	Min                        position         `yaml:"-" json:"min"`
	Name                       string           `yaml:"-" json:"name"`
	Planets                    map[int]Planet   `yaml:"planets" json:"-"`
	Position                   position         `yaml:"-" json:"position"`
	Radius                     float64          `yaml:"radius" json:"radius"`
	Regional                   bool             `yaml:"regional" json:"regional"`
	RegionID                   int              `yaml:"-" json:"region_id"`
	Security                   float64          `yaml:"security" json:"security"`
	SecurityClass              string           `yaml:"securityClass" json:"security_class"`
	SolarSystemID              int              `yaml:"solarSystemID" json:"solar_system_id"`
	SolarSystemNameID          int              `yaml:"solarSystemNameID" json:"solar_system_name_id"`
	Star                       StarType         `yaml:"star" json:"-"`
	StarID                     int              `yaml:"-" json:"star_id"`
	Stargates                  map[int]Stargate `yaml:"stargates" json:"-"`
	StargateIDs                []int            `yaml:"-" json:"stargates"`
	Stations                   []int            `yaml:"-" json:"stations"`
	SystemPlanets              []SystemPlanet   `yaml:"-" json:"planets"`
	SunTypeID                  int              `yaml:"sunTypeID" json:"sun_type_id"`
}

type AsteroidBelt struct {
	ID            int            `yaml:"-" json:"id"`
	Name          string         `yaml:"-" json:"name"`
	PositionArray []float64      `yaml:"position" json:"-"`
	Position      position       `yaml:"-" json:"position"`
	Statistics    StatisticsType `yaml:"statistics" json:"statistics"`
	SystemID      int            `yaml:"-" json:"system_id"`
	TypeID        int            `yaml:"typeID" json:"type_id"`
}

type StatisticsType struct {
	Density        float32 `yaml:"density" json:"density"`
	Eccentricity   float64 `yaml:"eccentricity" json:"eccentricity"`
	EscapeVelocity float64 `yaml:"escapeVelocity" json:"escape_velocity"`
	Fragmented     bool    `yaml:"fragmented" json:"fragmented"`
	Life           float64 `yaml:"life" json:"life"`
	Locked         bool    `yaml:"locked" json:"locked"`
	MassDust       float64 `yaml:"massDust" json:"mass_dust"`
	MassGas        float64 `yaml:"massGas" json:"mass_gas"`
	OrbitPeriod    float64 `yaml:"orbitPeriod" json:"orbit_period"`
	OrbitRadius    float64 `yaml:"orbitRadius" json:"orbit_radius"`
	Pressure       float64 `yaml:"pressure" json:"pressure"`
	Radius         float64 `yaml:"radius" json:"radius"`
	RotationRate   float64 `yaml:"rotationRate" json:"rotation_rate"`
	SpectralClass  string  `yaml:"spectralClass" json:"spectral_class"`
	SurfaceGravity float64 `yaml:"surfaceGravity" json:"surface_gravity"`
	Temperature    float64 `yaml:"temperature" json:"temperature"`
}

type StarStatisticsType struct {
	Age           int     `yaml:"age" json:"age"`
	Life          float64 `yaml:"life" json:"life"`
	Locked        bool    `yaml:"locked" json:"locked"`
	Luminosity    float64 `yaml:"luminosity" json:"luminosity"`
	Radius        float64 `yaml:"radius" json:"radius"`
	SpectralClass string  `yaml:"spectralClass" json:"spectral_class"`
	Temperature   int     `yaml:"temperature" json:"temperature"`
}

type Planet struct {
	AsteroidBelts     map[int]AsteroidBelt `yaml:"asteroidBelts" json:"-"`
	AsteroidBeltArray []int                `yaml:"-" json:"asteroid_belts"`
	CelestialIndex    int                  `yaml:"celestialIndex" json:"celestial_index"`
	Moons             map[int]Moon         `yaml:"moons" json:"-"`
	MoonArray         []int                `yaml:"-" json:"moons"`
	Name              string               `yaml:"-" json:"name"`
	PlanetAttributes  PlanetAttributesType `yaml:"planetAttributes" json:"planet_attributes"`
	PlanetID          int                  `yaml:"-" json:"planet_id"`
	PositionArray     []float64            `yaml:"position" json:"-"`
	Position          position             `yaml:"-" json:"position"`
	Radius            int                  `yaml:"radius" json:"radius"`
	Statistics        StatisticsType       `yaml:"statistics" json:"statistics"`
	SystemID          int                  `yaml:"-" json:"system_id"`
	TypeID            int                  `yaml:"typeID" json:"type_id"`
}

type PlanetAttributesType struct {
	HeightMap1   int  `yaml:"heightMap1" json:"height_map1"`
	HeightMap2   int  `yaml:"heightMap2" json:"height_map2"`
	Population   bool `yaml:"population" json:"population"`
	ShaderPreset int  `yaml:"shaderPreset" json:"shader_preset"`
}

type Moon struct {
	ID               int                  `yaml:"-" json:"moon_id"`
	Name             string               `yaml:"-" json:"name"`
	NPCStations      map[int]NPCStation   `yaml:"npcStations" json:"-"`
	PlanetAttributes PlanetAttributesType `yaml:"planetAttributes" json:"planet_attributes"`
	PositionArray    []float64            `yaml:"position" json:"-"`
	Position         position             `yaml:"-" json:"position"`
	Radius           int                  `yaml:"radius" json:"radius"`
	Stations         []int                `yaml:"-" json:"stations"`
	Statistics       StatisticsType       `yaml:"statistics" json:"statistics"`
	TypeID           int                  `yaml:"typeID" json:"type_id"`
}

type NPCStation struct {
	ConstellationID          int       `yaml:"constellationID" json:"constellation_id"`
	CorporationID            int       `yaml:"corporationID" json:"corporation_id"`
	DockingCostPerVolume     int       `yaml:"dockingCostPerVolume" json:"docking_cost_per_volume"`
	GraphicID                int       `yaml:"graphicID" json:"graphic_id"`
	IsConquerable            bool      `yaml:"isConquerable" json:"is_conquerable"`
	MaxShipVolumeDockable    int       `yaml:"maxShipVolumeDockable" json:"max_ship_volume_dockable"`
	OfficeRentalCost         int       `yaml:"officeRentalCost" json:"office_rental_cost"`
	OperationID              int       `yaml:"operationID" json:"operation_id"`
	OwnerID                  int       `yaml:"ownerID" json:"ownerID"`
	PositionArray            []float64 `yaml:"position" json:"-"`
	Position                 position  `yaml:"-" json:"position"`
	RegionID                 int       `yaml:"regionID" json:"region_id"`
	ReprocessingEfficiency   float32   `yaml:"reprocessingEfficiency" json:"reprocessing_efficiency"`
	ReprocessingHangarFlag   int       `yaml:"reprocessingHangarFlag" json:"reprocessing_hangar_flag"`
	ReprocessingStationsTake float32   `yaml:"reprocessingStationsTake" json:"reprocessing_stations_take"`
	Security                 float32   `yaml:"security" json:"security"`
	SolarSystemID            int       `yaml:"solarSystemID" json:"system_id"`
	StationID                int       `yaml:"stationID" json:"station_id"`
	StationName              string    `yaml:"stationName" json:"name"`
	StationTypeID            int       `yaml:"stationTypeID" json:"station_type_id"`
	TypeID                   int       `yaml:"typeID" json:"type_id"`
	UseOperationName         bool      `yaml:"useOperationName" json:"use_operation_name"`
}

type position struct {
	X float64 `yaml:"-" json:"x"`
	Y float64 `yaml:"-" json:"y"`
	Z float64 `yaml:"-" json:"z"`
}

type SystemPlanet struct {
	AsteroidBelts []int `yaml:"-" json:"asteroid_belts"`
	Moons         []int `yaml:"-" json:"moons"`
	PlanetID      int   `yaml:"-" json:"planet_id"`
}

type StarType struct {
	Age           int                `yaml:"-" json:"age"`
	ID            int                `yaml:"id" json:"id"`
	Luminosity    float64            `yaml:"-" json:"luminosity"`
	Name          string             `yaml:"-" json:"name"`
	Radius        int                `yaml:"radius" json:"radius"`
	SolarSystemID int                `yaml:"-" json:"solar_system_id"`
	SpectralClass string             `yaml:"-" json:"spectral_class"`
	Statistics    StarStatisticsType `yaml:"statistics" json:"statistics"`
	Temperature   int                `yaml:"-" json:"temperature"`
	TypeID        int                `yaml:"typeID" json:"type_id"`
}

type Stargate struct {
	Destination   int       `yaml:"destination" json:"destination"`
	ID            int       `yaml:"-" json:"id"`
	Name          string    `yaml:"-" json:"name"`
	PositionArray []float64 `yaml:"position" json:"-"`
	Position      position  `yaml:"-" json:"position"`
	TypeID        int       `yaml:"typeID" json:"type_id"`
}
