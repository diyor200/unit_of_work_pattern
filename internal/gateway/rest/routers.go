package rest

import "github.com/gin-gonic/gin"

func (h *Handler) NewRouter(router *gin.Engine) *gin.Engine {

	router.Handle("POST", "/user", h.AddUser)

	return router

}
