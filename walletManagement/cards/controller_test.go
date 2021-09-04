package cards

import (
	"errors"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type userServiceMock struct {
	getCardsMock   func() ([]Card, error)
	createCardMock func(card Card)
}

func (mock userServiceMock) getCards() ([]Card, error) {
	return mock.getCardsMock()
}

func (mock userServiceMock) createCards(card Card) {
}

func TestHandleGetCards(test *testing.T) {

	test.Run("should return 200 on successful getting list of cards", func(t *testing.T) {
		// GIVEN
		serviceMock := userServiceMock{}
		cardList := []Card{
			{Name: "test"},
		}
		httpRecorder := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/cards", HandleGetCards)

		serviceMock.getCardsMock = func() ([]Card, error) {
			return cardList, nil
		}
		UserService = serviceMock

		// restore mock
		oldService := UserService
		UserService = serviceMock
		defer func() {
			UserService = oldService
		}()

		// WHEN
		request, err := http.NewRequest(http.MethodGet, "/cards", nil)
		router.ServeHTTP(httpRecorder, request)

		// THEN
		assert.NoError(test, err)
		assert.Equal(test, 200, httpRecorder.Code)
		assert.Contains(test, httpRecorder.Body.String(), `test`)
	})

	test.Run("should return 500 on unexpected error", func(t *testing.T) {
		// GIVEN
		serviceMock := userServiceMock{}
		httpRecorder := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/cards", HandleGetCards)

		serviceMock.getCardsMock = func() ([]Card, error) {
			return nil, errors.New("failed to fetch from db")
		}

		// restore mock
		oldService := UserService
		UserService = serviceMock
		defer func() {
			UserService = oldService
		}()

		// WHEN
		request, _ := http.NewRequest(http.MethodGet, "/cards", nil)
		router.ServeHTTP(httpRecorder, request)

		// THEN
		assert.Equal(test, 500, httpRecorder.Code)
		assert.Contains(test, httpRecorder.Body.String(), `{"data":"Internal server error"}`)
	})

}
