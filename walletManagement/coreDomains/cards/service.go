package cards

type userService interface {
	getCards() ([]Card, error)
	createCards(card Card) error
	deleteCard(id string) (int64, error)
}

type userServiceImpl struct{}

var (
	UserService userService = userServiceImpl{}
)

func (userService userServiceImpl) getCards() ([]Card, error) {
	cards, err := UserRepo.getCards()
	return cards, err
}

func (userService userServiceImpl) deleteCard(id string) (int64, error) {
	deleteCount, err := UserRepo.deleteCard(id)
	return deleteCount, err
}

func (userService userServiceImpl) createCards(card Card) error {
	newCard := Card{
		Name:         card.Name,
		WalletType:   card.WalletType,
		WalletId:     card.WalletId,
		DailyLimit:   card.DailyLimit,
		MonthlyLimit: card.MonthlyLimit,
	}

	err := UserRepo.createCard(newCard)
	return err
}
