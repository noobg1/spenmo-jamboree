package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type userRepoMock struct {
	getCardsRepoMock func() ([]Card, error)
	createCardMock   func(card Card)
}

func (mock userRepoMock) getCards() ([]Card, error) {
	return mock.getCardsRepoMock()
}

func (mock userRepoMock) createCard(card Card) {
}

func TestGetCards(test *testing.T) {
	// GIVEN
	repoMock := userRepoMock{}
	cardList := []Card{
		{Name: "test"},
	}
	repoMock.getCardsRepoMock = func() ([]Card, error) {
		return cardList, nil
	}
	UserRepo = repoMock

	// WHEN
	result, _ := UserService.getCards()

	// THEN
	assert.Equal(test, result, cardList)
}

func TestCreateCards(test *testing.T) {
	// GIVEN
	repoMock := userRepoMock{}
	card := Card{
		Name: "test",
	}
	repoMock.createCardMock = func(card Card) {
	}

	UserRepo = repoMock

	// WHEN
	UserService.createCards(card)

	// THEN
	// TODO figure out to how to spy on methods
}
