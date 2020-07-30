package services

import (
	"brawlstars-api-test/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseURL = "https://api.brawlstars.com/v1/"

func GetBrawlStarsPlayer(player string) (*model.BrawlStarsPlayer, error) {
	fmt.Println(player)
	url := baseURL + "players/" + player
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("authorization", "Bearer "+"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjdlNGZjODc3LWE3MmEtNDhhZi1iYWY4LWIwOWFlNDgwZjkzMiIsImlhdCI6MTU5NjExMjk2Niwic3ViIjoiZGV2ZWxvcGVyLzYzZGE2MDI1LTU1MzktNzQzYy1iYWFiLWI5ZmExMzIzOThmZSIsInNjb3BlcyI6WyJicmF3bHN0YXJzIl0sImxpbWl0cyI6W3sidGllciI6ImRldmVsb3Blci9zaWx2ZXIiLCJ0eXBlIjoidGhyb3R0bGluZyJ9LHsiY2lkcnMiOlsiNjAuMzMuNzUuMTMwIl0sInR5cGUiOiJjbGllbnQifV19.5xHoau2muMlAF4u6GcBpHMCd_7tfuuIV2HFWkxv0Vsx-ax_j0qsB9hoS0ayYTcUntEQRjF_czykUiAU8TpMyCw")
	request.Header.Set("Accept", "application/json")
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	var brawlStarsPlayer model.BrawlStarsPlayer
	fmt.Println(string(data))
	err = json.Unmarshal(data, &brawlStarsPlayer)
	if err != nil {
		return nil, err
	}
	return &brawlStarsPlayer, nil
}
