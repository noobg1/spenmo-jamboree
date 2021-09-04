package cards

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spenmo-jamboree/walletManagement/common"
)

const (
	BaseRoute string = "/cards"
)

// getCards responds with the list of all cards as JSON.
// @Success 200 {array} Card
// @Router /cards/ [get]
func HandleGetCards(context *gin.Context) {
	cards, err := UserService.getCards()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"data": common.INTERNAL_SERVER_ERROR})
	} else {
		context.JSON(http.StatusOK, cards)
	}
}

// createCard responds with 201 on successful creation of card
// @Param data body Card{Name} true "Input Card name only"
// @Success 201
// @Router /cards/ [post]
func HandleCreateCards(context *gin.Context) {
	var card Card

	if err := context.BindJSON(&card); err != nil {
		return
	}

	UserService.createCards(card)

	context.JSON(http.StatusNoContent, gin.H{"data": "Ok"})
}
