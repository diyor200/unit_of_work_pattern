package v1

import "github.com/gin-gonic/gin"

func (h *Server) RegisterRoutes(router *gin.Engine) *gin.Engine {

	router.Handle("POST", "/user", h.AddUser)

	return router

}
