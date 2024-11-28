package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("Get", "/cafe?count=5&city=moscow", nil)

	fmt.Println(req)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	fmt.Println(list)

	assert.Equal(t, len(list), totalCount)
}

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("Get", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	assert.Equal(t, http.StatusOK, status)
}

func TestСityThatDoesntExist(t *testing.T) {
	cities := "moscow"

	req := httptest.NewRequest("Get", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	city := req.URL.Query().Get("city")

	assert.Equal(t, city, cities)
}
