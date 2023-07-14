package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"servers_api/controllers/requests"
	"servers_api/initializers"
	"servers_api/models"
	"servers_api/repositories"
	"servers_api/usecases"
	"testing"
)

var servers = []models.ServerModel{
	{
		Model: gorm.Model{
			ID: 1,
		},
		Name:   "ESGI",
		Type:   "MEDIUM",
		Status: "RUNNING",
	},
	{
		Model: gorm.Model{
			ID: 2,
		},
		Name: "McPvP", Type: "LARGE", Status: "STOPPED"},
	{
		Model: gorm.Model{
			ID: 3,
		},
		Name: "Hypixel", Type: "LARGE", Status: "RUNNING"},
	{
		Model: gorm.Model{
			ID: 4,
		},
		Name: "Stuga", Type: "SMALL", Status: "STARTED"},
}

func setupSuite(tb testing.TB) (func(tb testing.TB), *gin.Engine) {
	serverRepository := repositories.InMemoryServerRepository{Servers: servers}

	createServerUseCase := usecases.CreateServerUseCase{ServerRepository: serverRepository}
	findAllServersUseCase := usecases.FindAllServersUseCase{ServerRepository: serverRepository}
	findServerByIDUseCase := usecases.FindServerByIDUseCase{ServerRepository: serverRepository}
	serverController := ServerController{
		CreateServerUseCase:   createServerUseCase,
		FindAllServersUseCase: findAllServersUseCase,
		FindServerByIDUseCase: findServerByIDUseCase,
	}

	router := initializers.SetUpRouter()
	api := router.Group("/api/v1")
	api.POST("/servers", serverController.CreateServer)
	api.GET("/servers", serverController.FindAllServers)
	api.GET("/servers/:id", serverController.FindServerByID)

	return func(tb testing.TB) {
		log.Println("teardown suite")
	}, router
}

func setupCreateServerTest(tb testing.TB, router *gin.Engine, createServer requests.CreateServerRequest) (func(tb testing.TB), *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	body, _ := json.Marshal(createServer)
	request := httptest.NewRequest("POST", "/api/v1/servers", bytes.NewBuffer(body))
	router.ServeHTTP(recorder, request)

	return func(tb testing.TB) {
		log.Println("teardown test")
	}, recorder
}

func TestCreateServerHandler(t *testing.T) {
	tearDownSuite, router := setupSuite(t)
	defer tearDownSuite(t)

	table := []struct {
		name         string
		input        requests.CreateServerRequest
		expectedCode int
	}{
		{"should create a server", requests.CreateServerRequest{Name: "Server test", Type: "SMALL", Status: "RUNNING"}, http.StatusCreated},
		{"should return 400 if name is empty", requests.CreateServerRequest{Name: "", Type: "SMALL", Status: "RUNNING"}, http.StatusBadRequest},
		{"should return 400 if type is empty", requests.CreateServerRequest{Name: "Server test", Type: "", Status: "RUNNING"}, http.StatusBadRequest},
		{"should return 400 if status is empty", requests.CreateServerRequest{Name: "Server test", Type: "SMALL", Status: ""}, http.StatusBadRequest},
		{"should return 400 if type is invalid", requests.CreateServerRequest{Name: "Server test", Type: "INVALID", Status: "RUNNING"}, http.StatusBadRequest},
		{"should return 400 if status is invalid", requests.CreateServerRequest{Name: "Server test", Type: "SMALL", Status: "INVALID"}, http.StatusBadRequest},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			tearDownTest, recorder := setupCreateServerTest(t, router, tt.input)
			defer tearDownTest(t)

			assert.Equal(t, tt.expectedCode, recorder.Code)
		})
	}
}

func setupFindAllServersTest(tb testing.TB, router *gin.Engine) (func(tb testing.TB), *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/v1/servers", nil)
	router.ServeHTTP(recorder, request)

	return func(tb testing.TB) {
		log.Println("teardown test")
	}, recorder
}

func TestFindAllServersHandler(t *testing.T) {
	tearDownSuite, router := setupSuite(t)
	defer tearDownSuite(t)

	table := []struct {
		name         string
		expectedCode int
	}{
		{"should return 200", http.StatusOK},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			tearDownTest, recorder := setupFindAllServersTest(t, router)
			defer tearDownTest(t)

			assert.Equal(t, tt.expectedCode, recorder.Code)
		})
	}
}

func setupFindServerByIDTest(tb testing.TB, router *gin.Engine, id int) (func(tb testing.TB), *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/servers/%d", id), nil)
	router.ServeHTTP(recorder, request)

	return func(tb testing.TB) {
		log.Println("teardown test")
	}, recorder
}

func TestFindServerByIDHandler(t *testing.T) {
	tearDownSuite, router := setupSuite(t)
	defer tearDownSuite(t)

	table := []struct {
		name         string
		id           int
		expectedCode int
	}{
		{"should return 200 if server exists", 1, http.StatusOK},
		{"should return 404 when server does not exist", 5, http.StatusNotFound},
		{"should return 400 when id is invalid", -1, http.StatusBadRequest},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			tearDownTest, recorder := setupFindServerByIDTest(t, router, tt.id)
			defer tearDownTest(t)

			assert.Equal(t, tt.expectedCode, recorder.Code)
		})
	}
}
