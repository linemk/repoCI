package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"repoCI/mocks"
	"testing"
)

func TestBookHandler_GetAllBooks(t *testing.T) {
	mocker := new(mocks.GetBook)
	expectedBooks := []Book{
		{ID: 1, Name: "Book One"},
		{ID: 2, Name: "Book Two"},
	}

	mocker.On("GetAllBooks", mock.Anything, mock.Anything).Return(expectedBooks, nil)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	mocker.GetAllBooks(w, req)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)

	mocker.AssertExpectations(t)
}
