package model

type BrawlStarsPlayer struct {
	Tag                                  string               `json:"tag"`
	Name                                 string               `json:"name"`
	Icon                                 BrawlStarsPlayerIcon `json:"id"`
	Trophies                             int                  `json:"trophies"`
	HighestTrophies                      int                  `json:"highest_trophies"`
	PowerPlayPoints                      int                  `json:"powerPlayPoints"`
	ExpLevel                             int                  `json:"expLevel"`
	ExpPoints                            int                  `json:"expPoints"`
	IsQualifiedFromChampionshipChallenge bool                 `json:"isQualifiedFromChampionshipChallenge"`
	Brawlers                             []Brawler            `json: "brawlers"`
}

type BrawlStarsPlayerIcon struct {
	Id int `json:"id"`
}
