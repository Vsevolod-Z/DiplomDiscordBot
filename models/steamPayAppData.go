package models

type SteamPayAppData struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	NumInStock  int    `json:"num_in_stock"`
	Activation  string `json:"activation"`
	IsAvailable bool   `json:"is_available"`
	Image       string `json:"image"`
	Prices      struct {
		Rub int `json:"rub"`
	} `json:"prices"`
}
