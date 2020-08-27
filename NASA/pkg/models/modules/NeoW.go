package modules

//NeoLookUpRequest Request for the NEO API lookup
type NeoLookUpRequest struct {
	AsteroidID string `json:"asteroid_id"`
	Prefix     string `json:"prefix"`
}

//NeoWFeedRequest structure for the feed request
type NeoWFeedRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Prefix    string `json:"prefix"`
}

//NeoWFeedResponse general response from the NEO API
type NeoWFeedResponse struct {
	Links            NewWFeedLinks                `json:"links"`
	ElementCount     int                          `json:"element_count"`
	NearEarthObjects map[string][]NearEarthObject `json:"near_earth_objects"`
}

//NewWFeedLinks support to the link json structure
type NewWFeedLinks struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
	Self string `json:"self"`
}

//NeoWBrowseRequest Browser request structure
type NeoWBrowseRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pagesize"`
	Prefix   string `json:"prefix"`
}
type NeoWBroseResponse struct {
	Links            NewWFeedLinks     `json:"links"`
	PageInformation  PageInformation   `json:"page"`
	NearEarthObjects []NearEarthObject `json:"near_earth_objects"`
}

type PageInformation struct {
	Size          int `json:"size"`
	TotalElements int `json:"total_elements"`
	TotalPages    int `json:"total_pages"`
	Number        int `json:"number"`
}

type NearEarthObject struct {
	Links                        NewWFeedLinks       `json:"links"`
	ID                           string              `json:"id"`
	NeoReferenceID               string              `json:"neo_reference_id"`
	Name                         string              `json:"name"`
	NasaJPLUrl                   string              `json:"nasa_jpl_url"`
	AbsoluteMagnitudeH           float64             `json:"absolute_magnitude_h"`
	EstimatedDiameter            EstimatedDiameter   `json:"estimated_diameter"`
	IsPotentialHazardousAsteroid bool                `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData            []CloseApproachData `json:"close_approach_data"`
	IsSentryObject               bool                `json:"is_sentry_object"`
}

type EstimatedDiameter struct {
	Kilometers EstimatedDiameterThreshold `json:"kilometers"`
	Meters     EstimatedDiameterThreshold `json:"meters"`
	Miles      EstimatedDiameterThreshold `json:"miles"`
	Feet       EstimatedDiameterThreshold `json:"feet"`
}

type EstimatedDiameterThreshold struct {
	Min float64 `json:"estimated_diameter_min"`
	Max float64 `json:"estimated_diameter_max"`
}

type CloseApproachData struct {
	CloseApporachDate      string           `json:"close_approach_date"`
	CloseApporachDateFull  string           `json:"close_approach_date_full"`
	EpochDateCloseApproach int64            `json:"epoch_date_close_approach"`
	RelativeVelocity       RelativeVelocity `json:"relative_velocity"`
	MissDistance           MissDistance     `json:"miss_distance"`
	OrbitingBody           string           `json:"orbiting_body"`
}

type RelativeVelocity struct {
	KilometersPerSecond string `json:"kilometers_per_second"`
	KilometersPerHour   string `json:"kilometers_per_hour"`
	MilesPerHour        string `json:"miles_per_hour"`
}

type MissDistance struct {
	Astronomical string `json:"astronomical"`
	Lunar        string `json:"lunar"`
	Kilometers   string `json:"kilometers"`
	Miles        string `json:"miles"`
}
