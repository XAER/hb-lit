package models

type LoginBodyRequest struct {
	ApiKey string `json:"apiKey"`
	Pin    string `json:"pin"`
}

type LoginResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Token string `json:"token"`
}

type SeriesResponse struct {
	Status string       `json:"status"`
	Data   []SeriesData `json:"data"`
}

type SeriesData struct {
	Id                   int       `json:"id"`
	Name                 string    `json:"name"`
	Slug                 string    `json:"slug"`
	Image                string    `json:"image"`
	NameTranslations     []string  `json:"nameTranslations"`
	OverviewTranslations []string  `json:"overviewTranslations"`
	Aliases              []Aliases `json:"aliases"`
	FirstAired           string    `json:"firstAired"`
	LastAired            string    `json:"lastAired"`
	NextAired            string    `json:"nextAired"`
	Score                float64   `json:"score"`
	Status               Status    `json:"status"`
	OriginalCountry      string    `json:"originalCountry"`
	OriginalLanguage     string    `json:"originalLanguage"`
	DefaultSeasonType    int       `json:"defaultSeasonType"`
	IsOrderRandomized    bool      `json:"isOrderRandomized"`
	LastUpdated          string    `json:"lastUpdated"`
	AverageRuntime       int       `json:"averageRuntime"`
	Episodes             []Episode `json:"episodes"`
	Overview             string    `json:"overview"`
	Year                 string    `json:"year"`
}

type Aliases struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}

type Status struct {
	// Everything in the Status struct is nullable
	Id          int    `json:"id"`
	Name        string `json:"name"`
	RecordType  string `json:"recordType"`
	KeepUpdated bool   `json:"keepUpdated"`
}

type Episode struct {
	Id int `json:"id"`
}
