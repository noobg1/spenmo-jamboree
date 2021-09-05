package cards

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type userRepoMock struct {
	getCardsRepoMock func() ([]Card, error)
	createCardMock   func(card Card) error
	deleteCardMock   func(id string) (int64, error)
}

func (mock userRepoMock) getCards() ([]Card, error) {
	return mock.getCardsRepoMock()
}

func (mock userRepoMock) createCard(card Card) error {
	return mock.createCardMock(card)
}

func (mock userRepoMock) deleteCard(id string) (int64, error) {
	if id == "123" {
		var result int64 = 1
		return result, nil
	} else if id == "idToFail" {
		return 0, errors.New("failed to delete")
	}
	return 0, nil
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
	// restore mock
	oldRepo := UserRepo
	UserRepo = repoMock
	defer func() {
		UserRepo = oldRepo
	}()

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
	repoMock.createCardMock = func(card Card) error {
		return nil
	}

	// restore mock
	oldRepo := UserRepo
	UserRepo = repoMock
	defer func() {
		UserRepo = oldRepo
	}()

	// WHEN
	UserService.createCards(card)

	// THEN
	// TODO figure out to how to spy on methods
}

func TestDeleteCard(test *testing.T) {

	test.Run("successfully delete the record", func(t *testing.T) {
		// GIVEN
		repoMock := userRepoMock{}

		// restore mock
		oldRepo := UserRepo
		UserRepo = repoMock
		defer func() {
			UserRepo = oldRepo
		}()

		// WHEN
		result, _ := UserService.deleteCard("123")

		// THEN
		var expectedValue int64 = 1
		assert.Equal(test, result, expectedValue)
	})

	test.Run("if no record exists returns 0", func(t *testing.T) {
		// GIVEN
		repoMock := userRepoMock{}

		// restore mock
		oldRepo := UserRepo
		UserRepo = repoMock
		defer func() {
			UserRepo = oldRepo
		}()

		// WHEN
		result, _ := UserService.deleteCard("1234")

		// THEN
		var expectedValue int64 = 0
		assert.Equal(test, result, expectedValue)
	})

	test.Run("should 0 if no records are deleted", func(t *testing.T) {
		// GIVEN
		repoMock := userRepoMock{}

		// restore mock
		oldRepo := UserRepo
		UserRepo = repoMock
		defer func() {
			UserRepo = oldRepo
		}()

		// WHEN
		result, _ := UserService.deleteCard("1234")

		// THEN
		var expectedValue int64 = 0
		assert.Equal(test, result, expectedValue)
	})

	test.Run("should return error if db delete fails", func(t *testing.T) {
		// GIVEN
		repoMock := userRepoMock{}

		// restore mock
		oldRepo := UserRepo
		UserRepo = repoMock
		defer func() {
			UserRepo = oldRepo
		}()

		// WHEN
		_, error := UserService.deleteCard("idToFail")

		// THEN
		assert.Equal(test, error.Error(), "failed to delete")
	})
}
