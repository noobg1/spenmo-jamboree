package cards

import (
	"context"
	"log"
	"time"

	"github.com/spenmo-jamboree/walletManagement/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func getCardsRepo() []*Card {

	collection := utils.DbConnection.Database("jamboree").Collection("cards")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var cards []*Card
	for cur.Next(ctx) {
		var card Card
		err := cur.Decode(&card)
		if err != nil {
			log.Fatal(err)
		}

		cards = append(cards, &card)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return cards
}
