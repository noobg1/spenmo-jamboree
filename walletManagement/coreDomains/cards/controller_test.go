package cards

import (
	"errors"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type userServiceMock struct {
	getCardsMock   func() ([]Card, error)
	createCardMock func(card Card) error
	deleteCardMock func(id string) (int64, error)
}

func (mock userServiceMock) getCards() ([]Card, error) {
	return mock.getCardsMock()
}

func (mock userServiceMock) createCards(card Card) error {
	return nil
}

func (mock userServiceMock) deleteCard(id string) (int64, error) {
	if id == "123" {
		var result int64 = 1
		return result, nil
	} else if id == "idToFail" {
		return 0, errors.New("failed to delete")
	}
	return 0, nil
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

	test.Run("should return 500 on unexpected error if db calls fails for unexpected reason", func(t *testing.T) {
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
		assert.Contains(test, httpRecorder.Body.String(), `error`)
	})

}

func TestHandleCreateCards(test *testing.T) {
	test.Run("should return 204 on successful card creation", func(t *testing.T) {
		// GIVEN
		serviceMock := userServiceMock{}
		httpRecorder := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/cards", HandleCreateCards)

		serviceMock.createCardMock = func(card Card) error {
			return nil
		}

		// restore mock
		oldService := UserService
		UserService = serviceMock
		defer func() {
			UserService = oldService
		}()

		// WHEN
		body := `{"name": "jee", "walletType": "team", "walletId": "6134895c65039b0e7bd841bc", "dailyLimit": 120, "monthlyLimit": 240  }`
		request, _ := http.NewRequest(http.MethodPost, "/cards", strings.NewReader(body))
		router.ServeHTTP(httpRecorder, request)

		// THEN
		assert.Equal(test, 204, httpRecorder.Code)
	})

	test.Run("should return 400 on body payload validation failure", func(t *testing.T) {
		// GIVEN
		serviceMock := userServiceMock{}
		httpRecorder := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/cards", HandleCreateCards)

		serviceMock.createCardMock = func(card Card) error {
			return nil
		}

		// restore mock
		oldService := UserService
		UserService = serviceMock
		defer func() {
			UserService = oldService
		}()

		// WHEN
		body := `{"name": "jee", "walletType": "team", "walletId": "6134895c65039b0e7bd841bc", "monthlyLimit": 240  }`
		request, _ := http.NewRequest(http.MethodPost, "/cards", strings.NewReader(body))
		router.ServeHTTP(httpRecorder, request)

		// THEN
		assert.Equal(test, 400, httpRecorder.Code)
		assert.Contains(test, httpRecorder.Body.String(), `DailyLimit: zero value`)
	})
}

func TestHandleDeleteCard(test *testing.T) {
	test.Run("should return 200 on successful card deletion", func(t *testing.T) {
		// GIVEN
		serviceMock := userServiceMock{}
		httpRecorder := httptest.NewRecorder()
		router := gin.Default()
		router.DELETE("/cards/:id", HandleDeleteCard)

		// restore mock
		oldService := UserService
		UserService = serviceMock
		defer func() {
			UserService = oldService
		}()

		// WHEN
		request, _ := http.NewRequest(http.MethodDelete, "/cards/123", nil)
		router.ServeHTTP(httpRecorder, request)

		// THEN
		assert.Equal(test, 200, httpRecorder.Code)
		assert.Contains(test, httpRecorder.Body.String(), `{"data":1}`)
	})

	test.Run("should return 500 on something unexpected", func(t *testing.T) {
		// GIVEN
		serviceMock := userServiceMock{}
		httpRecorder := httptest.NewRecorder()
		router := gin.Default()
		router.DELETE("/cards/:id", HandleDeleteCard)

		// restore mock
		oldService := UserService
		UserService = serviceMock
		defer func() {
			UserService = oldService
		}()

		// WHEN
		request, _ := http.NewRequest(http.MethodDelete, "/cards/idToFail", nil)
		router.ServeHTTP(httpRecorder, request)

		// THEN
		assert.Equal(test, 500, httpRecorder.Code)
	})
}
