package user

import (
	"WB-test-L0/internal/delivery/api"
	"WB-test-L0/internal/domain/model"
	"WB-test-L0/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	ParamUUID = "uuid"

	GetPath  = "/entity/:" + ParamUUID
	PostPath = "/order"
)

type handler struct {
	userService service.Service
}

func NewHandler(service service.Service) api.Handler {
	return &handler{userService: service}
}

func (h *handler) NoRoute(c *gin.Context) {
	c.Status(http.StatusNotFound)
}

func (h *handler) Register(router *gin.Engine) {
	router.NoRoute(h.NoRoute)

	router.GET(GetPath, h.GetEntity)
	router.POST(PostPath, h.PostEntity)
}

func (h *handler) PostEntity(c *gin.Context) {
	var entity model.Entity
	err := c.ShouldBindJSON(&entity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(entity)
	message, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.userService.SetToNats(message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (h *handler) GetEntity(c *gin.Context) {
	entity, err := h.userService.FindByUUID(c.Param(ParamUUID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_uid":          entity.OrderUID,
		"track_number":       entity.TrackNumber,
		"entry":              entity.Entry,
		"delivery":           entity.Delivery,
		"payment":            entity.Payment,
		"items":              entity.Items,
		"locale":             entity.Locale,
		"internal_signature": entity.InternalSignature,
		"customer_id":        entity.CustomerID,
		"delivery_service":   entity.DeliveryService,
		"shardkey":           entity.ShardKey,
		"sm_id":              entity.SmID,
		"date_created":       entity.DateCreated,
		"oof_shard":          entity.OofShard,
	})
}
