package models

type SteamBuyData struct {
	Goods map[string]interface{} `json:"data"`
}
type SteamBuyAppDataWithoutPrice struct {
	IdGood    int    `json:"id_good"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Available int    `json:"available"` // 1-true , 0-false

}
type SteamBuyAppData struct {
	IdGood    int    `json:"id_good"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Available int    `json:"available"` // 1-true , 0-false
	Price     struct {
		Rub string `json:"rub"`
	} `json:"price"` // cant be array
}
