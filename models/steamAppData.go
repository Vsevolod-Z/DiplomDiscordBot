package models

import "encoding/json"

type SteamApp struct {
	Success bool         `json:"success"`
	Data    SteamAppData `json:"data"`
}

type SteamAppData struct {
	Type           string `json:"type"`
	Name           string `json:"name"`
	IsFree         bool   `json:"is_free"`
	AppID          int    `json:"steam_appid"`
	URL            string `json:"url"`
	DLC            []int  `json:"dlc"`
	Icon           string `json:"header_image"`
	PCRequirements struct {
		Minimum     []string `json:"minimum,omitempty"`
		Recommended []string `json:"recommended,omitempty"`
	} `json:"pc_requirements"`
	Developers         []string `json:"developers"`
	Publishers         []string `json:"publishers"`
	SteamPriceOverview struct {
		GameInitialPrice    int `json:"initial"`
		GameFinalPrice      int `json:"final"`
		GameDiscountPercent int `json:"discount_percent"`
	} `json:"price_overview"`
	Metacritic struct {
		Score int `json:"score"`
	} `json:"metacritic"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	Categories   []Category   `json:"categories"`
	Genres       []Genre      `json:"genres"`
	Screenshots  []Screenshot `json:"screenshots"`
	Movies       *[]Movie     `json:"movies,omitempty"`
	Achievements struct {
		Total       int           `json:"total"`
		Highlighted []Achievement `json:"highlighted"`
	} `json:"achievements"`
	ReleaseDate struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
	AppReview AppReview `json:"app_review"`
}
type Movie struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	MP4       struct {
		LQ string `json:"480"`
		HQ string `json:"max"`
	} `json:"mp4"`
}
type SteamRawAppData struct {
	Type                string          `json:"type"`
	Name                string          `json:"name"`
	IsFree              bool            `json:"is_free"`
	AppID               int             `json:"steam_appid"`
	URL                 string          `json:"url"`
	DLC                 []int           `json:"dlc"`
	Icon                string          `json:"header_image"`
	PCRequirementsSteam json.RawMessage `json:"pc_requirements,-"`
	PCRequirements      struct {
		Minimum     []string `json:"minimum,omitempty"`
		Recommended []string `json:"recommended,omitempty"`
	} `json:"pc_requirements_edited"`
	Developers         []string `json:"developers"`
	Publishers         []string `json:"publishers"`
	SteamPriceOverview struct {
		GameInitialPrice    int `json:"initial"`
		GameFinalPrice      int `json:"final"`
		GameDiscountPercent int `json:"discount_percent"`
	} `json:"price_overview"`
	Metacritic struct {
		Score int `json:"score"`
	} `json:"metacritic"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	Categories   []Category   `json:"categories"`
	Genres       []Genre      `json:"genres"`
	Screenshots  []Screenshot `json:"screenshots"`
	Movies       *[]Movie     `json:"movies,omitempty"`
	Achievements struct {
		Total       int           `json:"total"`
		Highlighted []Achievement `json:"highlighted"`
	} `json:"achievements"`
	ReleaseDate struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
}
type Screenshot struct {
	Id       int    `json:"id"`
	URL      string `json:"path_thumbnail"`
	PathFull string `json:"path_full"`
}
type Category struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
type Genre struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
type SteamListModel struct {
	Title  string `json:"title"`
	AppID  int    `json:"appid"`
	Price  int    `json:"price"`
	URL    string `json:"url"`
	Icon   string `json:"icon"`
	Rating int    `json:"rating"`
}

type Achievement struct {
	Name string `json:"name"`
	Icon string `json:"path"`
}

type SteamUser struct {
	Name           string `json:"personaname"`
	Avatar         string `json:"avatarfull"`
	UserLevel      int    `json:"player_level"`
	YearsOfService int    `json:"years_of_service"`
	Wishlist       []int  `json:"wishlist"`
}
type AppReview struct {
	Success      int          `json:"success"`
	QuerySummary QuerySummary `json:"query_summary"`
	Reviews      []Review     `json:"reviews"`
	Cursor       string       `json:"cursor"`
}
type QuerySummary struct {
	NumReviews      int    `json:"num_reviews"`
	ReviewScore     int    `json:"review_score"`
	ReviewScoreDesc string `json:"review_score_desc"`
	TotalPositive   int    `json:"total_positive"`
	TotalNegative   int    `json:"total_negative"`
	TotalReviews    int    `json:"total_reviews"`
}
type Review struct {
	RecommendationID         string `json:"recommendationid"`
	Author                   Author `json:"author"`
	Language                 string `json:"language"`
	Review                   string `json:"review"`
	TimestampCreated         int64  `json:"timestamp_created"`
	TimestampUpdated         int64  `json:"timestamp_updated"`
	VotedUp                  bool   `json:"voted_up"`
	VotesUp                  int    `json:"votes_up"`
	VotesFunny               int    `json:"votes_funny"`
	WeightedVoteScore        string `json:"weighted_vote_score"`
	CommentCount             int    `json:"comment_count"`
	SteamPurchase            bool   `json:"steam_purchase"`
	ReceivedForFree          bool   `json:"received_for_free"`
	WrittenDuringEarlyAccess bool   `json:"written_during_early_access"`
	HiddenInSteamChina       bool   `json:"hidden_in_steam_china"`
	SteamChinaLocation       string `json:"steam_china_location"`
}

type Author struct {
	SteamID              string `json:"steamid"`
	NumGamesOwned        int    `json:"num_games_owned"`
	NumReviews           int    `json:"num_reviews"`
	PlaytimeForever      int    `json:"playtime_forever"`
	PlaytimeLastTwoWeeks int    `json:"playtime_last_two_weeks"`
	PlaytimeAtReview     int    `json:"playtime_at_review"`
	LastPlayed           int64  `json:"last_played"`
}
