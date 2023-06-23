package models

type GOGData struct {
	Products []GOGAppData `json:"products"`
}

type GOGAppData struct {
	Price struct {
		BaseAmount  string `json:"baseAmount"`
		FinalAmount string `json:"finalAmount"`
		Discount    int    `json:"discount"`
	} `json:"price"`
	Title string `json:"title"`
	URL   string `json:"url"`
}
