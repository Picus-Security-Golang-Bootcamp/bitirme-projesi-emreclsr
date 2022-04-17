# Picus Final Project

## Introduction
This project is a sample e-commerce API that allows users to create new accounts, log-in, search products and fill 
their basket, create an order and cancel their orders. Admin can create new categories and products unlike the users. 

This project has been developed for Patika - Picus Security Go Bootcamp. 15 different tasks were completed and 
73 different tests were written. The project is designed according to RESTful principles. 

## Built with
* [Golang](https://go.dev/)
* [Gin](https://github.com/gin-gonic/gin)
* [Gorm](https://gorm.io/index.html)
* [PostgreSQL](postgresql.org)
* [JSON Web Tokens](https://jwt.io/)
* [Zap](https://github.com/uber-go/zap)
* [Swaggo-swag](https://github.com/swaggo/swag)

## Database Schema

<img src="https://cdn.discordapp.com/attachments/519918508998656028/965279531348930590/unknown.png" alt="Logo" width="100%">

## Usage
Code file runs on port 8000. After running the code, you can easily check each endpoint thanks to [swagger](https://swagger.io/) using the link below::

```
localhost:8000/swagger/index.html#/
```

### Sign Up
Users can sign up by sending a POST request to the following URL:
```
localhost:8000/user
```
Sample request:
```json
{
    "full_name": "testuser",
    "email": "test@test.com",
    "password": "test",
    "phone":"01234567898",
    "address": "test address",
    "role":"admin",
    "status":"active"
}
```
Password is hashed using bcrypt.
Role can be either `admin` or `customer`.

### Log-In
Users can log-in by sending a POST request to the following URL:
```
localhost:8000/login
```
Sample request:
```json
{
    "email": "test@test.com",
    "password": "test"
}
```
After successful login, the user has a JWT token for future requests. Token is valid for 1 hour.

### Create Category From CSV
Only admin can use this endpoint and create new categories, and product from CSV file. Endpoint is using with POST method.
```
localhost:8000/category
```
Sample CSV file:

|Type      |p_name   |p_price|p_stock|p_type    |
|----------|---------|-------|-------|----------|
|electronic|keyboard1|164    |23     |electronic|
|electronic|keyboard2|272    |20     |electronic|
|electronic|keyboard3|359    |22     |electronic|

### Get All Categories
Everyone can use this endpoint and get all categories. Endpoint is using with GET method.
```
localhost:8000/category
```

### Create / Update Basket
Users can update their basket by sending a PUT request to the following URL:
```
localhost:8000/basket
```
If the user has no basket, a new basket will be created with their User ID.
Sample request:
```json
{
    "user_id":1,
    "product_ids":[1,2,3],
    "amount":[3,2,4]
}
```
### Get Basket
Users can get their basket by sending a GET request to the following URL:
```
localhost:8000/basket
```

### Create an Order
Users can create an order by sending a POST request to the following URL:
```
localhost:8000/basket
```
In here, the stock control, maximum quantity of each product etc. will be checked. If everything is OK, the order will be created.

### Get Orders
Users can get their orders by sending a GET request to the following URL:
```
localhost:8000/order
```

### Cancel Order
Users can cancel their order by sending a PUT request to the following URL:
```
localhost:8000/order/:id
```
If order created less than 14 days ago, the order will be canceled.

### Create Product
Only admin can use this endpoint and create new products. Endpoint is using with POST method.
```
localhost:8000/product
```
Sample request:
```json
{
  "name": "new product",
  "price": 100,
  "stock": 10,
  "type": "electronic",
} 
```

### Get Products
Everyone can use this endpoint and get all products. Endpoint is using with GET method.
```
localhost:8000/product
```
Also in this endpoint using pagination. You can get products with given limit and sort options. Example:
```
localhost:8000/product?limit=10&page=10&sort=price
```
This endpoint will return last page of sorted products depends on their price. Page number starts from 0.


### Search Products
Everyone can use this endpoint and search products by given keyword. Endpoint is using with GET method.
```
localhost:8000/product/:word
```
This endpoint will return products that contains given keyword.

### Delete Product
Only admin can use this endpoint and delete products. Endpoint is using with DELETE method.
```
localhost:8000/product/:id
```
This endpoint actually not delete product from database, only change their deleted_at field.

### Update Product
Only admin can use this endpoint and update products. Endpoint is using with PUT method.
```
localhost:8000/product/:id
```
Sample request:
```json
{
  "name": "updated product",
  "price": 100,
  "stock": 10,
  "type": "electronic",
} 
```

## Coverage
In this project coverage tested with 73 different tests.

Code files and their coverage are listed below:

|               File                | Coverage |
|:---------------------------------:|:--------:|
| `authentication/loginhandler.go`  | `58.3%`  |
| `authentication/signuphandler.go` | `52.6%`  |
|     `authentication/token.go`     | `27.3%`  |
|        `basket/basket.go`         | `81.8%`  |
|        `basket/handler.go`        | `47.3%`  |
|      `basket/repository.go`       |  `100%`  |
|        `basket/service.go`        |  `100%`  |
|       `category/handler.go`       | `57.6%`  |
|     `category/repository.go`      | `76.5%`  |
|       `category/service.go`       | `69.2%`  |
|        `order/handler.go`         | `53.5%`  |
|         `order/order.go`          |  `100%`  |
|       `order/repository.go`       | `82.6%`  |
|        `order/service.go`         | `61.9%`  |
|       `product/handler.go`        | `42.7%`  |
|       `order/repository.go`       | `74.2%`  |
|        `order/service.go`         | `54.5%`  |
|       `user/repository.go`        | `75.0%`  |
|         `user/service.go`         |   `0%`   |
|          `user/user.go`           | `66.7%`  |





























