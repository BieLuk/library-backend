package book

import (
	"encoding/json"
	"fmt"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/service/books"
	"github.com/BieLuk/library-backend/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookController_CreateBook(t *testing.T) {
	// given
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	requestObject := dto.CreateBookRequest{
		Name:        "Test book name",
		Author:      "Test author name",
		ISBN:        "1234567890123",
		Description: utils.Pointer("test description"),
	}

	// when
	returnID := uuid.New()
	bookServiceMock := books.NewMockBookService(t)
	bookServiceMock.Mock.
		On("CreateBook", requestObject).
		Return(&dto.CreateBookResponse{ID: returnID}, nil).
		Once()

	c.Request = &http.Request{
		Header: make(http.Header),
	}
	utils.MockJsonPost(c, requestObject)

	// then
	bookController := NewBookController(bookServiceMock)
	bookController.CreateBook(c)

	response := fmt.Sprintf("{\"id\":\"%s\"}", returnID.String())
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response, w.Body.String())
}

func TestBookController_GetBook(t *testing.T) {
	// given
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	bookID := uuid.New()
	c.Params = gin.Params{
		{
			Key:   "id",
			Value: bookID.String(),
		},
	}
	responseObject := dto.GetBookResponse{
		ID:          bookID,
		Name:        "Test book name",
		Author:      "Test author name",
		ISBN:        "1234567890123",
		Description: utils.Pointer("test description"),
	}

	// when
	bookServiceMock := books.NewMockBookService(t)
	bookServiceMock.Mock.
		On("GetBook", bookID).
		Return(&responseObject, nil).
		Once()

	// then
	bookController := NewBookController(bookServiceMock)
	bookController.GetBook(c)

	responseJson, err := json.Marshal(responseObject)
	require.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(responseJson), w.Body.String())
}
