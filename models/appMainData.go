package models

/* type AppMainData struct {
	LocalID              int          `bson:"local_id"`
	Title                string       `bson:"title"`
	AppID                int          `bson:"appid"`
	Icon                 string       `bson:"icon"`
	Screenshots          []Screenshot `bson:"screenshots"`
	Categories           []Category   `bson:"categories"`
	Rating               int          `bson:"rating"`
	SteamPrice           int          `bson:"steam_price"`
	SteamDiscountPrecent int          `bson:"discount_percent"`
	SteamURL             string       `bson:"steam_url"`
	SteamPCRequirements  string       `bson:"pc_requirements"`
	SteamPayPrice        int          `bson:"steampay_price"`
	SteamPayURL          string       `bson:"steampay_url"`
	SteamBuyPrice        int          `bson:"steambuy_price"`
	SteamBuyURL          string       `bson:"steambuy_url"`
	GOGPrice             int          `bson:"gog_price"`
	GOGURL               string       `bson:"gog_url"`
} */

type AppMainData struct {
	LocalID         int             `bson:"local_id"`
	SteamAppData    SteamAppData    `bson:"steamData"`
	SteamBuyAppData SteamBuyAppData `bson:"steamBuyData"`
	SteamPayAppData SteamPayAppData `bson:"steamPayData"`
	GOGAppData      GOGAppData      `bson:"gogAppData"`
	PlayerPeeks     []PlayerPeek    `bson:"player_peeks"`
}
type PlayerPeek struct {
	TimePeriod string `bson:"time_period"`
	Count      string `bson:"count"`
}
