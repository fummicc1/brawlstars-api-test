package controller

import (
	"brawlstars-api-test/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlayer(context *gin.Context) {
	player := context.Param("player")
	brawlStarsPlayer, err := services.GetBrawlStarsPlayer(player)
	if err != nil {
		fmt.Println(err)
		responseError(context, err)
	}
	context.JSON(http.StatusOK, *brawlStarsPlayer)
}

func responseError(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, err.Error())
}
