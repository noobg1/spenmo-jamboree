package common

import (
	"os"
	"time"
)

var (
	DB_PASS_ENV                 = os.Getenv("MONGO_PASS")
	DB_CONNECTION_STRING string = "mongodb+srv://jamboree" + ":" + DB_PASS_ENV + "@cluster0.tkoku.mongodb.net/jamboree"
)

const (
	COLLECTION_NAME   string        = "cards"
	DB_NAME           string        = "jamboree"
	EXECUTION_TIMEOUT time.Duration = 30
)
