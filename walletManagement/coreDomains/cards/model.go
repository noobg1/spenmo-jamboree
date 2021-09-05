package cards

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" validate:"min=1,max=16,regexp=^[a-zA-Z]*$,nonnil"`
	WalletType   string             `json:"walletType" validate:"min=4,max=4,regexp=^[(user)|(team)]*$,nonnil"`
	WalletId     primitive.ObjectID `json:"walletId" validate:"min=12,max=12,nonnil"`
	DailyLimit   float64            `json:"dailyLimit" validate:"nonzero"`
	MonthlyLimit float64            `json:"monthlyLimit" validate:"nonzero"`
}
