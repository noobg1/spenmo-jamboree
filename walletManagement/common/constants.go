package common

import (
	"os"
	"time"
)

var (
	DB_PASS_ENV             string = os.Getenv("MONGO_PASS")
	DB_CONNECTION_STRING    string = "mongodb+srv://jamboree" + ":" + DB_PASS_ENV + "@cluster0.tkoku.mongodb.net/jamboree"
	REDIS_PASS_ENV          string = os.Getenv("REDIS_PASS")
	REDIS_CONNECTION_STRING string = "redis://default:" + REDIS_PASS_ENV + "@redis-19802.c252.ap-southeast-1-1.ec2.cloud.redislabs.com:19802/0"
)

const (
	COLLECTION_NAME   string        = "cards"
	DB_NAME           string        = "jamboree"
	EXECUTION_TIMEOUT time.Duration = 30
)
