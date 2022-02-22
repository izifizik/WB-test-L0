package user

import (
	"WB-test-L0/internal/delivery/api"
	"WB-test-L0/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	GetPath  = "/:uuid"
	PostPath = "/pub"
)

type handler struct {
	userService service.Service
}

func NewHandler(service service.Service) api.Handler {
	return &handler{userService: service}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(GetPath, h.GetEntity)
	router.POST(PostPath, h.PostEntity)
}

func (h *handler) GetEntity(c *gin.Context) {

}

func (h *handler) PostEntity(c *gin.Context) {

}
