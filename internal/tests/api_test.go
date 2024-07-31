package tests

import (
	"bytes"
	"e-commerce/internal/domain"
	"e-commerce/internal/handler"
	"e-commerce/internal/repository"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"
)

var testDB *gorm.DB

func setupProductRouter() *gin.Engine {
	r := gin.Default()
	productRepo := repository.NewProductRepository(testDB)
	ph := handler.NewProductHandler(productRepo)
	r.POST("/products", ph.CreateProduct)
	r.GET("/products", ph.GetAllProducts)
	r.GET("/products/:id", ph.GetProductByID)
	r.PUT("/products/:id", ph.UpdateProduct)
	r.DELETE("/products/:id", ph.DeleteProduct)
	r.GET("/products/search/name", ph.SearchProductsByName)
	r.GET("/products/search/category", ph.SearchProductsByCategory)
	return r
}

func setupDB() *gorm.DB {
	dsn := "user=testuser password=testpassword dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.Product{})
	return db
}

func TestMain(m *testing.M) {
	testDB = setupDB()
	code := m.Run()
	os.Exit(code)
}

func TestCreateProduct(t *testing.T) {
	setupDB()
	router := setupProductRouter()
	product := domain.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       10.0,
		Category:    "Test Category",
		Quantity:    5,
	}
	productJSON, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Product created successfully!", response["message"])
}

func TestGetAllProducts(t *testing.T) {
	setupDB()
	router := setupProductRouter()

	product := domain.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       10.0,
		Category:    "Test Category",
		Quantity:    5,
	}
	productJSON, _ := json.Marshal(product)

	req, _ := http.NewRequest("GET", "/products", bytes.NewBuffer(productJSON))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var products []domain.Product
	json.Unmarshal(w.Body.Bytes(), &products)
	assert.Greater(t, len(products), 0)
}

func TestGetProductByID(t *testing.T) {
	setupDB()
	router := setupProductRouter()
	product := setupTestProduct()

	req, _ := http.NewRequest("GET", "/products/"+strconv.Itoa(int(product.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetchedProduct domain.Product
	json.Unmarshal(w.Body.Bytes(), &fetchedProduct)
	assert.Equal(t, product.ID, fetchedProduct.ID)
}

func TestUpdateProduct(t *testing.T) {
	router := setupProductRouter()
	product := setupTestProduct()

	updatedProduct := domain.Product{
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       20.0,
		Category:    "Updated Category",
		Quantity:    10,
	}
	updatedProductJSON, _ := json.Marshal(updatedProduct)

	req, _ := http.NewRequest("PUT", "/products/"+strconv.Itoa(int(product.ID)), bytes.NewBuffer(updatedProductJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Product updated successfully!", response["message"])
}

func TestDeleteProduct(t *testing.T) {
	router := setupProductRouter()
	product := setupTestProduct()

	req, _ := http.NewRequest("DELETE", "/products/"+strconv.Itoa(int(product.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Product deleted successfully!", response["message"])
}

func TestSearchProductsByName(t *testing.T) {
	router := setupProductRouter()
	setupTestProduct()

	req, _ := http.NewRequest("GET", "/products/search/name?name=Test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var products []domain.Product
	json.Unmarshal(w.Body.Bytes(), &products)
	assert.Greater(t, len(products), 0)
}

func TestSearchProductsByCategory(t *testing.T) {
	router := setupProductRouter()
	setupTestProduct()

	req, _ := http.NewRequest("GET", "/products/search/category?category=Test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var products []domain.Product
	json.Unmarshal(w.Body.Bytes(), &products)
	assert.Greater(t, len(products), 0)
}

func setupTestProduct() domain.Product {
	product := domain.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       10.0,
		Category:    "Test Category",
		Quantity:    5,
	}

	testDB.Create(&product)
	return product
}

func setupRouter1(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	user := router.Group("/user")
	{
		user.POST("/", userHandler.CreateUser)
		user.GET("/", userHandler.GetAllUsers)
		user.GET("/:id", userHandler.GetUserByID)
		user.PUT("/:id", userHandler.UpdateUser)
		user.DELETE("/:id", userHandler.DeleteUser)
		user.GET("/search/:name", userHandler.SearchUsersByName)
		user.GET("/search/email/:email", userHandler.SearchUsersByEmail)
	}

	return router
}

func TestCreateUser1(t *testing.T) {
	dsn := "host=localhost user=testuser password=testpassword dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to the database: %v", err)
	}

	db.AutoMigrate(&domain.User{})

	router := setupRouter1(db)

	user := domain.User{
		Name:             "John Doe",
		Email:            "john.doel@example.com",
		Address:          "123 Elm Street",
		RegistrationDate: time.Now(),
		Role:             "client",
	}

	reqBody, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/user/", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message":"User created successfully!"}`, w.Body.String())
}

func TestGetAllUsers1(t *testing.T) {
	dsn := "host=localhost user=testuser password=testpassword dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to the database: %v", err)
	}

	db.AutoMigrate(&domain.User{})

	db.Create(&domain.User{
		Name:             "John Doe",
		Email:            "john.doe@example.com",
		Address:          "123 Elm Street",
		RegistrationDate: time.Now(),
		Role:             "client",
	})

	router := setupRouter1(db)

	req := httptest.NewRequest(http.MethodGet, "/user/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUser1(t *testing.T) {
	dsn := "host=localhost user=testuser password=testpassword dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to the database: %v", err)
	}

	db.AutoMigrate(&domain.User{})

	user := domain.User{
		Name:             "John Doe",
		Email:            "john.doe@example.com",
		Address:          "123 Elm Street",
		RegistrationDate: time.Now(),
		Role:             "client",
	}
	db.Create(&user)

	router := setupRouter1(db)

	updatedUser := domain.User{
		Name: "Jane Doe",
	}
	reqBody, _ := json.Marshal(updatedUser)
	req := httptest.NewRequest(http.MethodPut, "/user/"+strconv.Itoa(int(user.ID)), bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"User updated successfully!"}`, w.Body.String())
}

func TestDeleteUser1(t *testing.T) {
	dsn := "host=localhost user=testuser password=testpassword dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to the database: %v", err)
	}

	db.AutoMigrate(&domain.User{})

	user := domain.User{
		Name:             "John Doe",
		Email:            "john.doe@example.com",
		Address:          "123 Elm Street",
		RegistrationDate: time.Now(),
		Role:             "client",
	}
	db.Create(&user)

	router := setupRouter1(db)

	req := httptest.NewRequest(http.MethodDelete, "/user/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"User deleted successfully!"}`, w.Body.String())
}

func setupTestDB() (*gorm.DB, func()) {
	dsn := "user=testuser password=testpassword dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Order{}, &domain.User{}, &domain.Product{})

	cleanup := func() {
		db.Migrator().DropTable(&domain.Order{}, &domain.User{}, &domain.Product{})
	}

	return db, cleanup
}

func TestCreateOrder(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	orderHandler := handler.NewOrderHandler(orderRepo, userRepo, productRepo)

	user := domain.User{ID: 1}
	product := domain.Product{ID: 1}
	db.Create(&user)
	db.Create(&product)

	order := domain.Order{
		UserID:     1,
		ProductIDs: []uint{1},
		TotalPrice: 100.0,
		Status:     "new",
	}
	body, _ := json.Marshal(order)

	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	orderHandler.CreateOrder(c)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expectedMessage := "Order created successfully!"
	if response["message"] != expectedMessage {
		t.Errorf("handler returned unexpected body: got %v want %v", response["message"], expectedMessage)
	}
}

func TestGetOrderByID(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	orderHandler := handler.NewOrderHandler(orderRepo, userRepo, productRepo)

	user := domain.User{ID: 2}
	product := domain.Product{ID: 3}
	order := domain.Order{ID: 1, UserID: user.ID, ProductIDs: []uint{product.ID}, TotalPrice: 100.0, Status: "new"}
	db.Create(&user)
	db.Create(&product)
	db.Create(&order)

	req, _ := http.NewRequest(http.MethodGet, "/orders/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	orderHandler.GetOrderByID(c)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response domain.Order
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if response.ID != 1 {
		t.Errorf("handler returned unexpected body: got %v want %v", response.ID, 1)
	}
}

func TestGetAllOrders(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	orderHandler := handler.NewOrderHandler(orderRepo, userRepo, productRepo)

	user := domain.User{ID: 1}
	product := domain.Product{ID: 1}
	orders := []domain.Order{
		{ID: 1, UserID: 1, ProductIDs: []uint{1}, TotalPrice: 100.0, Status: "new"},
		{ID: 2, UserID: 1, ProductIDs: []uint{1}, TotalPrice: 200.0, Status: "processing"},
	}
	db.Create(&user)
	db.Create(&product)
	db.Create(&orders)

	req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	orderHandler.GetAllOrders(c)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response []domain.Order
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(response) != 2 {
		t.Errorf("handler returned unexpected number of orders: got %v want %v", len(response), 2)
	}
}
