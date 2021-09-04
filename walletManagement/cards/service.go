package cards

type userService interface {
	getCards() ([]Card, error)
	createCards(card Card)
}

type userServiceImpl struct{}

var (
	UserService userService = userServiceImpl{}
)

func (userService userServiceImpl) getCards() ([]Card, error) {
	cards, err := UserRepo.getCards()
	return cards, err
}

func (userService userServiceImpl) createCards(card Card) {
	newCard := Card{
		Name: card.Name,
	}

	UserRepo.createCard(newCard)
}
