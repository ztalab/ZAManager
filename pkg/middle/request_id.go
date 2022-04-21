package middle

import (
	"github.com/ztalab/ZAManager/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// RequestID adds a unique request id to the context
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.NewString()
		logger.NewContext(c, zap.String("requestID", reqID))
		c.Set("requestID", reqID)
		c.Next()
	}
}
