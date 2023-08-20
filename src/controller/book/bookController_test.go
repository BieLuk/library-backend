package book

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/service/books"
	"github.com/BieLuk/library-backend/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

	// when
	bookController := NewBookController(bookServiceMock)
	bookController.CreateBook(c)

	// then
	response := fmt.Sprintf("{\"id\":\"%s\"}", returnID.String())
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response, w.Body.String())
}
