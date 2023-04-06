package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HanderError_when_endpoints_returns_internal_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}
	handlerFunc := HandlerError(endpoint)

	println(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())

}

func Test_HanderError_when_endpoints_returns_domain_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("bad request error")
	}
	handlerFunc := HandlerError(endpoint)

	println(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "bad request error")

}

func Test_HanderError_when_endpoints_returns_object(t *testing.T) {
	assert := assert.New(t)

	type ObjectNaruto struct {
		ID   int
		Name string
	}

	myobject := ObjectNaruto{ID: 1, Name: "Nagato"}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return myobject, 201, nil
	}
	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	personNarutoAnime := ObjectNaruto{}
	fmt.Println(&personNarutoAnime)

	json.Unmarshal(res.Body.Bytes(), &personNarutoAnime)
	assert.Equal(myobject, personNarutoAnime)

}
