package modules

type TLECollectionRequest struct {
	Search   string `json:"search"`
	PM       string `json:"pm"`
	Sort     string `json:"sort"`
	SortDir  string `json:"sort-dir"`
	Page     int    `json:"page"`
	PageSize string `json:"page-size"`
	Prefix   string `json:"prefix"`
}

type TLERecordRequest struct {
	ID     int64  `json:"id"`
	Prefix string `json:"prefix"`
}

type TLERecordResponse struct {
	Record TLEMember `json:"record"`
}

type TLECollectionResponse struct {
	Context    string              `json:"@context"`
	ID         string              `json:"@id"`
	Type       string              `json:"@type"`
	TotalItems int64               `json:"totalItems"`
	Members    []TLEMember         `json:"member"`
	Parameters TLESearchParameters `json:"parameters"`
	View       TLEView             `json:"view"`
}

type TLEView struct {
	ID       string `json:"@id"`
	Type     string `json:"@type"`
	First    string `json:"first"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Last     string `json:"last"`
}

type TLEMember struct {
	ID          string `json:"@id"`
	Type        string `json:"@type"`
	SatelitteID int64  `json:"satelliteId"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
}

type TLESearchParameters struct {
	Search   string `json:"search"`
	Sort     string `json:"sort"`
	SortDir  string `json:"sort-dir"`
	Page     int    `json:"page"`
	PageSize int    `json:"page-size"`
}
