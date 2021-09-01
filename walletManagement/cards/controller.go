package cards

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getCards responds with the list of all albums as JSON.
// @Success 200 {array} album
// @Router /examples/groups/{group_id}/accounts/{account_id} [get]
func HandleGetCards(context *gin.Context) {
	cards := GetCardsService()
	context.IndentedJSON(http.StatusOK, cards)
}
