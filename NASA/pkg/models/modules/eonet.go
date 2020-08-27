package modules

type EonetRequest struct {
	Prefix     string `json:"prefix"`
	Source     string `json:"source"`
	Status     string `json:"status"`
	Limit      int    `json:"limit"`
	Days       int    `json:"days"`
	CategoryID int    `json:"category_id"`
}

type EonetEventResponse struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Link        string       `json:"link"`
	Events      []EonetEvent `json:"events"`
}

//Categoy Description

type EonetCategoryResponse struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Link        string          `json:"link"`
	Categories  []EonetCategory `json:"categories"`
}

//Event Descriptions

type EonetEvent struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Link        string          `json:"link"`
	Categories  []EonetCategory `json:"categories"`
	Sources     []EonetSource   `json:"sources"`
	Geometries  []EonetGeometry `json:"geometries"`
}

type EonetCategory struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Layers      string `json:"layers"`
}

type EonetSource struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type EonetGeometry struct {
	Date        string    `json:"date"`
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
