package hello_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"fogcms-server/internal/modules/hello"
)

func TestHelloWorldHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := hello.NewHandler()
	router := gin.New()
	router.GET("/", handler.HelloWorld)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "{\"message\":\"Hello World\"}"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHelloModule_RegisterRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	module := hello.NewModule()
	router := gin.New()
	module.RegisterRoutes(router)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Module route returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
