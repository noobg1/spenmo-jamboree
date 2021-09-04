package wallets

import (
	"github.com/spenmo-jamboree/walletManagement/coreDomains/cards"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserWallet struct {
	Cards []cards.Card       `json:"cards,omitempty" bson:"cards,omitempty"`
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name"`
}
