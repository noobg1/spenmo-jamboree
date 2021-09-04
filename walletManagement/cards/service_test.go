package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type userRepoMock struct {
	getCardsRepoMock func() []Card
	createCardMock   func(card Card)
}

func (mock userRepoMock) getCards() []Card {
	return mock.getCardsRepoMock()
}

func (mock userRepoMock) createCard(card Card) {
}

func TestGetCards(test *testing.T) {
	// GIVEN
	serviceMock := userRepoMock{}
	cardList := []Card{
		{Name: "test"},
	}
	serviceMock.getCardsRepoMock = func() []Card {
		return cardList
	}
	UserRepo = serviceMock

	// WHEN
	result := UserService.getCards()

	// THEN
	assert.Equal(test, result, cardList)
}

func TestCreateCards(test *testing.T) {
	// GIVEN
	serviceMock := userRepoMock{}
	card := Card{
		Name: "test",
	}
	serviceMock.createCardMock = func(card Card) {
	}

	UserRepo = serviceMock

	// WHEN
	UserService.createCards(card)

	// THEN
	// TODO figure out to how to spy on methods
}
