package modules

type DiscosRequest struct {
	Prefix  string   `json:"prefix"`
	Sort    string   `json:"sort"`
	Filter  string   `json:"filter"`
	Page    string   `json:"page"`
	Include string   `json:"include"`
	Fields  []string `json:"fields"`
}
type DiscosResponse struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			CosparID    string  `json:"cosparId"`
			Satno       int     `json:"satno"`
			VimpelID    int     `json:"vimpelId"`
			Name        string  `json:"name"`
			ObjectClass string  `json:"objectClass"`
			Mass        int     `json:"mass"`
			Shape       string  `json:"shape"`
			Length      float64 `json:"length"`
			Height      float64 `json:"height"`
			Depth       int     `json:"depth"`
			XSectMax    float64 `json:"xSectMax"`
			XSectMin    float64 `json:"xSectMin"`
			XSectAvg    float64 `json:"xSectAvg"`
		} `json:"attributes"`
		Relationships struct {
			States struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"states"`
			InitialOrbits struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"initialOrbits"`
			Operators struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"operators"`
			Launch struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"launch"`
			DestinationOrbits struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"destinationOrbits"`
			Reentry struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"reentry"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Links struct {
		Self  string      `json:"self"`
		First string      `json:"first"`
		Last  string      `json:"last"`
		Next  string      `json:"next"`
		Prev  interface{} `json:"prev"`
	} `json:"links"`
	Meta struct {
		Pagination struct {
			TotalPages  int `json:"totalPages"`
			CurrentPage int `json:"currentPage"`
			PageSize    int `json:"pageSize"`
		} `json:"pagination"`
	} `json:"meta"`
}
