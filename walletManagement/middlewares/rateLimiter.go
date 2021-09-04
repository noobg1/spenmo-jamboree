package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	libredis "github.com/go-redis/redis/v8"
	"github.com/spenmo-jamboree/walletManagement/common"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

const (
	RATE_LIMIT_FREQUENCY string = "10-M"
	// https://github.com/ulule/limiter#usage
)

func RateLimiter() gin.HandlerFunc {
	defer func() {
		if recoverResult := recover(); recoverResult != nil {
			log.Printf("Recovering from panic in printAllOperations error is: %v \n", recoverResult)
		}
	}()

	// Define a limit rate to 4 requests per hour.
	rate, err := limiter.NewRateFromFormatted(RATE_LIMIT_FREQUENCY)
	if err != nil {
		log.Panic(err)
	}

	// Create a redis client.
	option, err := libredis.ParseURL(common.REDIS_CONNECTION_STRING)
	if err != nil {
		log.Panic(err)
	}
	client := libredis.NewClient(option)

	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: "rate-limiter",
	})
	if err != nil {
		log.Panic(err)
	}

	// Create a new middleware with the limiter instance.
	middleware := mgin.NewMiddleware(limiter.New(store, rate))
	return middleware
}
