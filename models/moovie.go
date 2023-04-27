package models

type Moovies struct {
	Id       int      `json:"id"`
	Adult    bool     `json:"adult"`
	Path     string   `json:"backdrop_path"`
	Budget   float64  `json:"budget"`
	Genres   []Genres `json:"genres"`
	HomePage string   `json:"homepage"`
	Language string   `json:"original_language"`
	Title    string   `json:"original_name"`
}

type Genres struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
