package cards

import (
	"github.com/spenmo-jamboree/walletManagement/wallets"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	wallets.Wallet
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
}
