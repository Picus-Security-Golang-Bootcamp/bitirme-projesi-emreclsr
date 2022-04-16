package main

import (
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/basket"
	"github.com/emreclsr/picusfinal/category"
	"github.com/emreclsr/picusfinal/db"
	"github.com/emreclsr/picusfinal/handlers"
	"github.com/emreclsr/picusfinal/logger"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/repositories"
	services "github.com/emreclsr/picusfinal/services"
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
)

func main() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize global logger
	//logger.NewLogger()
	//defer logger.Close()
	logger.InitLogger()

	DB, err := db.Connect()
	if err != nil {
		zap.L().Fatal("Error connecting to database", zap.Error(err))
	}
	err = DB.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{}, &basket.Basket{}, &order.Order{})
	if err != nil {
		zap.L().Fatal("Error auto migrating database", zap.Error(err))
	}

	token := authentication.NewToken()
	repos := repositories.NewRepositories(DB)
	servs := services.NewServices(DB, *repos)
	hands := handlers.NewHandlers(*servs, token)

	r := gin.Default()
	r.POST("/user", hands.User.SignUp)                        // 1
	r.POST("/login", hands.Authentication.Login)              // 2
	r.POST("/category", hands.Category.CreateCategoryFromCSV) // 3
	r.GET("/category", hands.Category.GetAllCategories)       // 4
	r.PUT("/basket", hands.Basket.UpdateBasket)               // 5-7
	r.GET("/basket", hands.Basket.GetBasket)                  // 6
	r.POST("/basket", hands.Basket.CreateAnOrder)             // 8
	r.GET("/order", hands.Order.GetOrders)                    // 9
	r.PUT("/order/:id", hands.Order.CancelOrder)              // 10
	r.POST("/product", hands.Product.CreateProduct)           // 11
	r.GET("/product", hands.Product.GetProducts)              // 12
	r.GET("/product/:word", hands.Product.Search)             // 13
	r.DELETE("/product/:id", hands.Product.DeleteProduct)     // 14
	r.PUT("/product/:id", hands.Product.UpdateProduct)        // 15

	log.Fatal(r.Run(":8000"))

}
