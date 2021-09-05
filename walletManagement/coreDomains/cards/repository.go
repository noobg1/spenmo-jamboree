package cards

import (
	"context"
	"log"
	"time"

	"github.com/spenmo-jamboree/walletManagement/common"
	"github.com/spenmo-jamboree/walletManagement/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo interface {
	getCards() ([]Card, error)
	createCard(card Card) error
	deleteCard(id string) (int64, error)
}

type userRepoImpl struct{}

var (
	UserRepo userRepo = userRepoImpl{}
)

func collectionInit() (*mongo.Collection, context.Context, context.CancelFunc) {
	collection := utils.DbConnection.Database(common.DB_NAME).Collection(common.COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), common.EXECUTION_TIMEOUT*time.Second)

	return collection, ctx, cancel
}

func (userRepo userRepoImpl) createCard(card Card) error {
	collection, ctx, cancel := collectionInit()
	defer cancel()

	_, err := collection.InsertOne(ctx, card)

	if err != nil {
		log.Panicln(err)
	}
	return err
}

func (userRepo userRepoImpl) deleteCard(id string) (int64, error) {
	collection, ctx, cancel := collectionInit()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		log.Println(err)
	}
	return deleteResult.DeletedCount, err
}

func (userRepo userRepoImpl) getCards() ([]Card, error) {
	collection, ctx, cancel := collectionInit()
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Println(err)
	}
	defer cur.Close(ctx)
	var cards []Card
	for cur.Next(ctx) {
		var card Card
		err := cur.Decode(&card)
		if err != nil {
			log.Println(err)
		}

		cards = append(cards, card)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}

	return cards, err
}
