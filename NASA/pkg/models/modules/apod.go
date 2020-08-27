package modules

//ApodRequest structure
type ApodRequest struct {
	Prefix string `json:"prefix"`
	Date   string `json:"date"`
	HD     bool   `json:"hd"`
}

//ApodResponse response
type ApodResponse struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"serviceversion"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}
