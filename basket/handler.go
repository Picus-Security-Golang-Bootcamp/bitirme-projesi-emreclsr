package basket

import (
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BasketHandler struct {
	basketServ  BasketService
	token       authentication.Token
	userServ    user.UserService
	orderServ   order.OrderService
	productServ product.ProductService
}

func NewBasketHandler(bs BasketService, token authentication.Token, userServ user.UserService, orderServ order.OrderService, productServ product.ProductService) *BasketHandler {
	return &BasketHandler{
		basketServ:  bs,
		token:       token,
		userServ:    userServ,
		orderServ:   orderServ,
		productServ: productServ,
	}
}

func (h *BasketHandler) UpdateBasket(c *gin.Context) {
	claims, err := h.token.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized for this action"})
	}
	var basket Basket

	err = c.ShouldBindJSON(&basket)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// If basket not created yet, create one
	bskt, err := h.basketServ.GetByUserId(claims.UserID)
	if bskt.ID == 0 {
		err = h.basketServ.CreateBasket(claims.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	basket.UserID = claims.UserID

	for key, id := range basket.ProductIds {
		p, errs := h.productServ.Get(uint(id))
		if errs != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if int64(p.Stock)-basket.Amount[key] < 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Not enough stock from product with name: " + string(p.Name)})
			return
		}
		basket.Products = append(basket.Products, *p)
	}
	basket.CalculateTotalPrice()

	if err := h.basketServ.UpdateBasket(&basket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "item added successfully"})
}

func (h *BasketHandler) GetBasket(c *gin.Context) {
	claims, err := h.token.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized for this action"})
		return
	}
	// If basket not created yet, create one
	basket, err := h.basketServ.GetByUserId(claims.UserID)
	if basket.ID == 0 {
		//err = h.basketServ.CreateBasket(claims.UserID)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	return
		//}
		c.JSON(http.StatusOK, gin.H{"message": "basket is empty"})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(basket.ProductIds) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "basket is empty"})
		return
	}
	for _, id := range basket.ProductIds {
		p, errs := h.productServ.Get(uint(id))
		if errs != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		basket.Products = append(basket.Products, *p)
	}

	c.JSON(http.StatusOK, gin.H{"basket": basket})
}

func (h *BasketHandler) CreateAnOrder(c *gin.Context) {
	claims, err := h.token.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized for this action"})
		return
	}
	basket, err := h.basketServ.GetByUserId(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for key, id := range basket.ProductIds {
		p, errs := h.productServ.Get(uint(id))
		if errs != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p.Stock = p.Stock - int(basket.Amount[key])
		err := h.productServ.Update(p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		basket.Products = append(basket.Products, *p)
	}

	// Check permitted items count and amount in basket
	status, err := basket.CheckItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if status == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	o := basket.ToOrder()
	err = h.orderServ.Create(o)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	emptyBasket := &Basket{UserID: claims.UserID}
	if err := h.basketServ.UpdateBasket(emptyBasket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": o})
}
