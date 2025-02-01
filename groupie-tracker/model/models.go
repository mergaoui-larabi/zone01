package model

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate uint16   `json:"creationDate"`
	Firstalbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Location     []string `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	ConcertDate  []string `json:"dates"`
	Relations    string   `json:"relations"`
	Relation     relation `json"datesLocations"`
}

type relation struct {
}
