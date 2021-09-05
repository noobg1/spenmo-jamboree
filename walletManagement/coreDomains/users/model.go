package users

import (
	"github.com/spenmo-jamboree/walletManagement/coreDomains/wallets"
)

type User struct {
	wallets.UserWallet
	Id   string `json:"id"`
	Name string `json:"name"`
}
