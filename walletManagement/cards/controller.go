package cards

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getCards responds with the list of all albums as JSON.
// @Success 200 {array} album
// @Router /examples/groups/{group_id}/accounts/{account_id} [get]
func HandleGetCards(context *gin.Context) {
	cards := UserService.getCards()
	context.IndentedJSON(http.StatusOK, cards)
}

func HandleCreateCards(context *gin.Context) {
	var card Card

	if err := context.BindJSON(&card); err != nil {
		return
	}

	UserService.createCards(card)

	context.IndentedJSON(http.StatusOK, "ok")
}
