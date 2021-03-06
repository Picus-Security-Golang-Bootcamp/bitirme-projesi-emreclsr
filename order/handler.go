package order

import (
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderServ OrderService
	token     authentication.TokenInterface
}
type IOrderHandler interface {
	GetOrders(c *gin.Context)
	CancelOrder(c *gin.Context)
}

var _ IOrderHandler = &OrderHandler{}

func NewOrderHandler(orderServ OrderService, token authentication.TokenInterface) IOrderHandler {
	return &OrderHandler{orderServ, token}
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Security TokenJWT
// @Success 200 {object} Order
// @Failure 401 "Unauthorized"
// @Failure 500 "Server error"
// @Router /order [get]
func (h *OrderHandler) GetOrders(c *gin.Context) {
	zap.L().Info("GetOrders handler triggered")
	claim, err := h.token.VerifyToken(c)
	if err != nil {
		zap.L().Error("Token verification failed in get orders handler", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	orders, err := h.orderServ.List(claim.UserID)
	if err != nil {
		zap.L().Error("Error while getting orders handler", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// CancelOrder godoc
// @Summary Cancel order
// @Description Cancel order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Security TokenJWT
// @Success 200 "Order canceled"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 500 "Server error"
// @Router /order/{id} [put]
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	zap.L().Info("CancelOrder handler triggered")
	claim, err := h.token.VerifyToken(c)
	if err != nil {
		zap.L().Error("Error while verifying token in cancel order handler", zap.Error(err))
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
		zap.L().Error("Error while getting order in cancel order handler", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if order.UserID != claim.UserID {
		zap.L().Error("User is not authorized to cancel order", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to cancel this order"})
		return
	}
	if !order.CheckTime() {
		zap.L().Error("Order can not be canceled", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can not cancel this order because it is not in the time limit"})
		return
	}
	order.IsCanceled = true

	if err := h.orderServ.Update(order); err != nil {
		zap.L().Error("Error while updating order in cancel order handler", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Order canceled"})
}
