//
//  Simple Solr
//
//  Copyright © 2020. All rights reserved.
//

package api_test

import (
	"github.com/moemoe89/go-solr-edo/api/api_struct/form"
	"github.com/moemoe89/go-solr-edo/api/mocks"
	"github.com/moemoe89/go-solr-edo/config"
	"github.com/moemoe89/go-solr-edo/routers"

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessCreate(t *testing.T) {
	userForm := form.UserForm{
		ID:       1,
		Name:    "name",
		Address: "address",
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Create", userForm).Return(http.StatusCreated, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/create", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFailCreate(t *testing.T) {
	userForm := form.UserForm{
		ID:       1,
		Name:    "name",
		Address: "address",
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Create", userForm).Return(http.StatusInternalServerError, errors.New("Unexpected error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/create", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFailValidationCreate(t *testing.T) {
	userForm := &form.UserForm{
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/create", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestSuccessSelect(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Select", "", "").Return(nil, 0, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/select", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFailSelect(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Select", "", "").Return(nil, http.StatusInternalServerError, errors.New("Unexpected error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/select", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSuccessDelete(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Delete", "1").Return(0, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/delete/1", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFailDelete(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Delete", "1").Return(http.StatusInternalServerError, errors.New("Unexpected error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/delete/1", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
