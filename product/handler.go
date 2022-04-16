package product

import (
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/pagination"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductService ProductService
	token          authentication.TokenInterface
}
type IProductHandler interface {
	CreateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	Search(c *gin.Context)
	DeleteProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
}

var _ IProductHandler = &ProductHandler{}

func NewProductHandler(service ProductService, token authentication.TokenInterface) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
		token:          token,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	zap.L().Info("CreateProduct handler triggered")
	claim, err := h.token.VerifyToken(c)
	if err != nil {
		zap.L().Error("Token verification failed in create product handler", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}
	if claim.Role != "admin" {
		zap.L().Error("Non admin user tried to create product", zap.String("user", claim.Role))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		zap.L().Error("Error in binding json in create product handler", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.ProductService.Create(&product); err != nil {
		zap.L().Error("Error in creating product (handler)", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

//func (h *ProductHandler) GetProducts(c *gin.Context) {
//	zap.L().Info("GetProducts handler triggered")
//	_, err := h.token.VerifyToken(c)
//	if err != nil {
//		zap.L().Error("Token verification failed in get products handler", zap.Error(err))
//		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//		return
//	}
//	products, err := h.ProductService.List()
//	if err != nil {
//		zap.L().Error("Error in getting products (handler)", zap.Error(err))
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, products)
//}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	pg := pagination.GeneratePaginationRequest(c)
	pg.URLPath = c.Request.URL.Path
	pagi, err := h.ProductService.List(pg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pagi)
}

func (h *ProductHandler) Search(c *gin.Context) {
	zap.L().Info("Search product handler triggered")
	word := c.Param("word")
	products, err := h.ProductService.Search(word)
	if err != nil {
		zap.L().Error("Error in searching product (handler)", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	zap.L().Info("DeleteProduct handler triggered")
	claim, err := h.token.VerifyToken(c)
	if err != nil {
		zap.L().Error("Token verification failed in delete product handler", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}
	if claim.Role != "admin" {
		zap.L().Error("Non admin user tried to delete product", zap.String("user", claim.Role))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.ProductService.Delete(uint(id)); err != nil {
		zap.L().Error("Error in deleting product (handler)", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	zap.L().Info("UpdateProduct handler triggered")
	claim, err := h.token.VerifyToken(c)
	if err != nil {
		zap.L().Error("Token verification failed in update product handler", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}
	if claim.Role != "admin" {
		zap.L().Error("Non admin user tried to update product", zap.String("user", claim.Role))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		zap.L().Error("Error in binding json to product (handler)", zap.Error(err))
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)
	if err := h.ProductService.Update(&product); err != nil {
		zap.L().Error("Error in updating product (handler)", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
