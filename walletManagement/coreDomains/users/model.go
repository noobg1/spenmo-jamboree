package users

import "github.com/spenmo-jamboree/walletManagement/coreDomains/cards"

type User struct {
	cards.Card
	id   string `json:"id"`
	name string `json:"name"`
}
