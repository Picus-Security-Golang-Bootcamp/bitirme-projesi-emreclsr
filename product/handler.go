package product

import (
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/pagination"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

var validate *validator.Validate
var _ IProductHandler = &ProductHandler{}

func NewProductHandler(service ProductService, token authentication.TokenInterface) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
		token:          token,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param body body Product true "Product"
// @Security TokenJWT
// @Success 201 {object} Product
// @Failure 400 "Bad request"
// @Failure 401 "You are not authorized to perform this action"
// @Failure 500 "Server error"
// @Router /product [post]
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
	validate = validator.New()
	err = validate.Struct(&product)
	if err != nil {
		zap.L().Error("Error in json which sended data to create product handler", zap.Error(err))
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

// GetProducts godoc
// @Summary Get products with pagination
// @Description Get products with pagination
// @Tags Product
// @Accept  json
// @Produce  json
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Param sort query string true "sort"
// @Success 200 {object} pagination.Pagination
// @Failure 500 "Server error"
// @Router /product [get]
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

// Search godoc
// @Summary Search products
// @Description Search products
// @Tags Product
// @Accept  json
// @Produce  json
// @Param word path string true "Search word"
// @Success 200 {object} Product
// @Failure 500 "Server error"
// @Router /product/{word} [get]
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

// DeleteProduct godoc
// @Summary Delete products
// @Description Delete products
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Security TokenJWT
// @Success 200 "Product deleted"
// @failure 400 "Bad request"
// @Failure 401 "You are not authorized to perform this action"
// @Failure 500 "Server error"
// @Router /product/{id} [delete]
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

// UpdateProduct godoc
// @Summary Update products
// @Description Update products
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param body body Product true "Product"
// @Security TokenJWT
// @Success 200 {object} Product
// @failure 400 "Bad request"
// @Failure 401 "You are not authorized to perform this action"
// @Failure 422 "Unprocessable Entity"
// @Failure 500 "Server error"
// @Router /product/{id} [put]
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
	validate = validator.New()
	err = validate.Struct(&product)
	if err != nil {
		zap.L().Error("Error in json which sended data to create product handler", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.ProductService.Update(&product); err != nil {
		zap.L().Error("Error in updating product (handler)", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
