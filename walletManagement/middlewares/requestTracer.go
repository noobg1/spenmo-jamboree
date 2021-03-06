package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spenmo-jamboree/walletManagement/common"
)

func Tracer() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestTenant := context.GetHeader(common.TRACER_ID)
		if requestTenant == "" {
			requestTenant = uuid.New().String()
		}
		context.Set(common.TRACER_ID, requestTenant)
		context.Next()
	}
}
