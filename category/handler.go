package category

import (
	"encoding/csv"
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/product"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	set "github.com/mpvl/unique"
	"go.uber.org/zap"
	"net/http"
	"sort"
	"strconv"
)

type CategoryHandler struct {
	catServ     CategoryService
	token       authentication.TokenInterface
	productServ product.ProductService
}
type ICategoryHandler interface {
	GetAllCategories(c *gin.Context)
	CreateCategoryFromCSV(c *gin.Context)
}

var _ ICategoryHandler = &CategoryHandler{}
var validate *validator.Validate

func NewCategoryHandler(catServ CategoryService, token authentication.TokenInterface, productServ product.ProductService) ICategoryHandler {
	return &CategoryHandler{
		catServ:     catServ,
		token:       token,
		productServ: productServ,
	}
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags Category
// @Accept  json
// @Produce  json
// @Success 200 {object} Category
// @Failure 500 "Server error"
// @Router /category [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	zap.L().Info("GetAllCategories handler triggered")
	categories, err := h.catServ.List()
	if err != nil {
		zap.L().Error("Error while getting categories (handler)", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategoryFromCSV godoc
// @Summary Create categories from csv
// @Description Create categories from csv
// @Tags Category
// @Accept  json
// @Produce  json
// @Param csv formData file true "csv"
// @Security TokenJWT
// @Success 200 "Categories created successfully"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 500 "Server error"
// @Router /category [post]
func (h *CategoryHandler) CreateCategoryFromCSV(c *gin.Context) {
	zap.L().Info("CreateCategoryFromCSV handler triggered")
	claim, err := h.token.VerifyToken(c)
	if err != nil {
		zap.L().Error("Error while verifying token in create category from csv", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if claim.Role != "admin" {
		zap.L().Error("Non admin user tried to create category from csv", zap.String("user", claim.Role))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}
	csvPartFile, _, openErr := c.Request.FormFile("csv")
	if openErr != nil {
		zap.L().Error("Error while opening csv file", zap.Error(openErr))
		c.JSON(http.StatusBadRequest, gin.H{"error": openErr.Error()})
		return
	}
	csvLines, readErr := csv.NewReader(csvPartFile).ReadAll()
	if readErr != nil {
		zap.L().Error("Error while reading csv file", zap.Error(readErr))
		c.JSON(http.StatusBadRequest, gin.H{"error": readErr.Error()})
		return
	}
	var catList []string
	var product product.Product

	for i, each := range csvLines {
		if i > 0 {
			catList = append(catList, each[0])
			product.Name = each[1]
			p, err := strconv.ParseFloat(each[2], 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			product.Price = p
			ps, err := strconv.Atoi(each[3])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			product.Stock = ps
			product.Type = each[4]
			validate = validator.New()
			err = validate.Struct(&product)
			if err != nil {
				zap.L().Error("Error in json which sended data to create product handler", zap.Error(err))
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := h.productServ.Create(&product); err != nil {
				zap.L().Error("Error while creating product (handler)", zap.Error(err))
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			product.ID = 0
		}
	}

	sort.Strings(catList)
	set.Strings(&catList)
	for _, i := range catList {
		category := Category{
			Type: i,
		}
		validate = validator.New()
		err = validate.Struct(&category)
		if err != nil {
			zap.L().Error("Error in json which sended data to create product handler", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := h.catServ.Create(&category); err != nil {
			zap.L().Error("Error while creating category (handler)", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categories created successfully"})
}
