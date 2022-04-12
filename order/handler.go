package order

import (
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderServ OrderService
	token     authentication.Token
}

func NewOrderHandler(orderServ OrderService, token authentication.Token) *OrderHandler {
	return &OrderHandler{orderServ, token}
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	claim, err := h.token.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	usrId := claim.UserID
	orders, err := h.orderServ.List(usrId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) CancelOrder(c *gin.Context) {
	claim, err := h.token.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	orderId := c.Param("id")
	oid, err := strconv.Atoi(orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := h.orderServ.Get(uint(oid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if order.UserID != claim.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to cancel this order"})
		return
	}
	if !order.CheckTime() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can not cancel this order because it is not in the time limit"})
		return
	}
	order.IsCanceled = true

	if err := h.orderServ.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Order canceled"})
}
