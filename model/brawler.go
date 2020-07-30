package model

type Brawler struct {
	Id              int         `json:"id"`
	Name            string      `json:"name"`
	Power           int         `json:"power"`
	Rank            int         `json:"rank"`
	Trophies        int         `json:"trophies"`
	HighestTrophies int         `json:"highestTrophies"`
	StarPowers      []StarPower `json:"starPowers"`
	Gadgets         []Gadget    `json:"gadgets"`
}

type Gadget struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type StarPower struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
