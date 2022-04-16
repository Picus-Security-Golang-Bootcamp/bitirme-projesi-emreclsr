package db

import (
	"fmt"
	"github.com/emreclsr/picusfinal/user"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DBTestConnect() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_TEST_HOST"),
		os.Getenv("DB_TEST_PORT"),
		os.Getenv("DB_TEST_USER"),
		os.Getenv("DB_TEST_PASSWORD"),
		os.Getenv("DB_TEST_NAME"))

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Fatal("Failed to get database connection", zap.Error(err))
		return nil, fmt.Errorf("failed to get database connection: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		zap.L().Error("Failed to ping database", zap.Error(err))
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}
	//err = db.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{}, &basket.Basket{}, &order.Order{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	return db, nil
}

var Hastoken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlJvbGUiOiJhZG1pbiIsImV4cCI6NTI1MDA1MDQ1MX0.cAx3260XwE3dRkRFyM_VEOSs7n1AbmZPbV9QPeeNwoE"

func AddUser(db *gorm.DB) {
	var usr user.User
	usr.Email = "test@test.com"
	usr.Password = "test"
	serv := user.NewUserService(user.NewUserRepository(db))
	err := serv.Create(&usr)
	if err != nil {
		fmt.Println(err)
	}

}

func DropDB(db *gorm.DB) {
	//db.Raw("DROP table users")
	//db.DropTable(user.User{})
	db.Exec("DROP schema public cascade; CREATE SCHEMA public;")
}
