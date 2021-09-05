package cards

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spenmo-jamboree/walletManagement/common"
	"gopkg.in/validator.v2"
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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		context.JSON(http.StatusOK, cards)
	}
}

// deleteCard responds with 200 on successful deletion of card
// @Param id path string true
// @Success 200
// @Router /cards/:id [delete]
func HandleDeleteCard(context *gin.Context) {
	id := context.Params.ByName("id")
	deleteCount, err := UserService.deleteCard(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": deleteCount})
	}
}

// createCard responds with 201 on successful creation of card
// @Param data body Card{Name} true "Input Card name only"
// @Success 201
// @Router /cards/ [post]
func HandleCreateCards(context *gin.Context) {
	var card Card

	if err := context.BindJSON(&card); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validations
	if err := validator.Validate(card); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if card.MonthlyLimit < card.DailyLimit {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": common.MONTHLY_LIMIT_IS_LOWER_THAN_DAILY,
		})
		return
	}

	// service call
	err := UserService.createCards(card)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err, // TODO map to AppErrors
		})
		return
	}

	// response
	context.JSON(http.StatusNoContent, gin.H{"data": "Ok"})
}
