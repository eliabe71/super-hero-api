//ok
package types
type Powerstats struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

type Biography struct {
	FullName        string   `json:"full-name"`
	AlterEgos       string   `json:"alter-egos"`
	Aliases         []string `json:"aliases"`
	PlaceOfBirth    string   `json:"place-of-birth"`
	FirstAppearance string   `json:"first-appearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}

type Appearance struct {
	Gender    string   `json:"gender"`
	Race      string   `json:"race"`
	Height    []string `json:"height"`
	Weight    []string `json:"weight"`
	Eyecolor  string   `json:"eye-color"`
	HairColor string   `json:"hair-color"`
}
type Work struct {
	Occupation 		string `json:"occupation"`
	BaseOfOperation string `json:"base"`
}
type Connections struct {
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}
type Image struct {
	Url string `json:"url"`
}
type Response struct{
	Response  string   `json:"response"`
}

/// Tipo auxiliar para requisiçao do json em conections
type NameAux struct{
	NameAux     string `json:"name"`
}
type SuperAndVilan struct {
	Name             string 	`json:"name"`
	Uuid        	 string     `json:"id"`
	Powerstats  	 Powerstats  `json:"powerstats""`
	Biography   	 Biography   `json:"biography"`
	Appearance 	 	 Appearance  `json:"appearance"`
	Work          	 Work        `json:"work"`
	Connections   	 Connections `json:"connections"`
	Image          	 Image       `json:"image"`
}

