package reacthandler

import "github.com/gin-gonic/gin"

func (h *Handler) GinHandler() gin.HandlerFunc {
	return gin.WrapF(h.HandleStatic)
}
