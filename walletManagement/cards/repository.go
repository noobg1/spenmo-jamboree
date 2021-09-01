package cards

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spenmo-jamboree/walletManagement/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func getCardsRepo() []Card {

	collection := utils.DbConnection.Database("jamboree").Collection("cards")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	cards := []Card{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		// card := Card{Id: result[0]._id, Name: result[1].name}
		for _, value := range result {
			fmt.Println(value.Value)

		}
		// append(cards, card)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return cards
}
