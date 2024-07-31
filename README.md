# Go - Service for managing an online store

### REST API for managing an online store that includes several microservices: Users, Products, Orders, Payments and API Gateway. The service support CRUD operations, published on GitHub, embedded in Render, and run using Docker Compose and Makefile.

## Getting Started

### Prerequisites:
    - Go 1.22 or later
    - Docker
### Installation:
1. Clone the repository:
   ```bash
   https://github.com/Aytya/Store-server-HL
   ```
2. Navigate into the project directory:
   ```bash
    cd Store-server-HL
   ```
3. Install dependencies:
   ```bash
    go get -u "github.com/gin-gonic/gin"
    go get -u "github.com/lib/pq"
    go get -u "github.com/joho/godotenv"
    go get -u "github.com/spf13/viper"
    go get -u "github.com/stripe/stripe-go"
    go get -u "github.com/go-playground/validator/v10"
    go get -u "github.com/stretchr/testify/assert"
    go get -u "gorm.io/driver/postgres"
   ```

##  Build and Run Locally:
### Build the application:
   ```bash
   make build
   ```
### Run the application:
   ```bash
   make run
   ```
### Stop the application:
   ```bash
   make down
   ```

## API Endpoints:
### User:
#### Create a New User:
- URL: http://localhost:8080/user
- Method: POST
- Request Body:
 ```bash
    {
      "name": "John Doe",
      "email": "john.doe@example.com",
      "role": "USER"
    }
 ```

#### Update an Existing User:
- URL: http://localhost:8080/user/:id
- Method: PUT
- Request Body:
 ```bash
    {
      "name": "John Doel",
      "email": "john.doel@example.com",
      "role": "USER"
    }
 ```

#### Get All Users:
- URL: http://localhost:8080/user
- Method: GET

#### Delete an Existing User:
- URL: http://localhost:8080/user/:id
- Method: DELETE

#### Get User By Id
- URL: http://localhost:8080/user/:id
- Method: GET

#### Search an Existing User By Name
- URL: http://localhost:8080/user/search/:name
- Method: GET

#### Search an Existing User By Email
- URL: http://localhost:8080/user/search/email/:email
- Method: GET

### Product: 
#### Create a New Product:
- URL: http://localhost:8080/products
- Method: POST
- Request Body:
 ```bash
    {
        "name": "Product 1",
        "description": "Description of Product 1",
        "price": 19.99,
        "category": "Category A"
    }
 ```

#### Update an Existing Product:
- URL: http://localhost:8080/products/:id
- Method: PUT
- Request Body:
 ```bash
    {
        "name": "Product 2",
        "description": "Description of Product 1",
        "price": 21.99,
        "category": "Category A"
    }
 ```

#### Get All Products:
   - URL: http://localhost:8080/products
   - Method: GET

#### Get Product by ID:
   - URL: http://localhost:8080/products/:id
   - Method: GET

#### Search Products by Name:
 - URL: http://localhost:8080/products/search/:name
 - Method: GET

#### Search Products by Category:
- URL: http://localhost:8080/products/search/category/:category
- Method: GET

### Order:
#### Create a New Order:
- URL: http://localhost:8080/orders
- Method: POST
- Request Body:
 ```bash
    {
        "user_id": 1,
        "product_ids": [1, 2],
        "total_price": 69.97,
        "order_date": "2024-07-06T12:00:00Z",
        "status": "new"
    }
 ```

#### Update an Existing Order:
- URL: http://localhost:8080/orders/:id
- Method: PUT
- Request Body:
 ```bash
    {
       "user_id": 1,
       "product_ids": [1, 2, 3],
       "total_price": 89.96,
       "order_date": "2024-07-06T12:00:00Z",
       "status": "shipped"
    }
 ```

#### Get All Orders:
- URL: http://localhost:8080/orders
- Method: GET

#### Get Order by ID:
- URL: http://localhost:8080/orders/:id
- Method: GET

#### Search Orders by Status:
- URL: http://localhost:8080/orders/search
- Method: GET

#### Search Orders by User ID:
- URL: http://localhost:8080/orders/search/:user
- Method: GET

### Payment:
#### Create a New Payment:
- URL: http://localhost:8080/payments
- Method: POST
- Request Body:
 ```bash
     {
          "user_id": 1,
          "order_id": 1,
          "amount": 69.97,
    }
 ```

#### Update an Existing Payment:
- URL: http://localhost:8080/payments/:id
- Method: PUT
- Request Body:
 ```bash
     {
          "user_id": 1,
          "order_id": 1,
          "amount": 79.97
     }
 ```

#### Get All Payments:
- URL: http://localhost:8080/payments
- Method: GET

#### Get Payment by ID:
- URL: http://localhost:8080/payments/:id
- Method: GET

#### Search Payments by User ID:
- URL: http://localhost:8080/payments/search/user/:user_id
- Method: GET

#### Search Payments by Order ID:
- URL: http://localhost:8080/payments/search/:order_id
- Method: GET

#### Search Payments by Status:
- URL: http://localhost:8080/payments/search
- Method: GET

### Swagger Documentation
- URL: http://localhost:8080/swagger/index.html#/
