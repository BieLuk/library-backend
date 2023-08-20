package borrow

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/apperr"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/service/borrows"
	"github.com/BieLuk/library-backend/src/validate"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type BorrowController interface {
	CreateBorrow(c *gin.Context)
	IsBookBorrowed(c *gin.Context)
	ReturnBorrowedBook(c *gin.Context)
}

type borrowController struct {
	borrowService borrows.BorrowService
}

func NewBorrowController(borrowService borrows.BorrowService) *borrowController {
	return &borrowController{
		borrowService: borrowService,
	}
}

// CreateBorrow creates model.Borrow object in database
func (bc *borrowController) CreateBorrow(c *gin.Context) {
	var request dto.CreateBorrowRequest
	if err := validate.BindAndValidateAny(c, &request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			apperr.NewAppErr(apperr.BAD_REQUEST, fmt.Sprintf("cannot unmarshall request object: %v", err)))
		return
	}

	response, err := bc.borrowService.CreateBorrow(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot create borrow: %v", err)))
		return
	}

	c.JSON(http.StatusOK, response)
}

// IsBookBorrowed checks if book with given ID is borrowed
func (bc *borrowController) IsBookBorrowed(c *gin.Context) {
	bookID := uuid.MustParse(c.Param("id"))
	response, err := bc.borrowService.IsBookBorrowed(bookID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot check if books is borrowed: %v", err)))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (bc *borrowController) ReturnBorrowedBook(c *gin.Context) {
	var request dto.ReturnBorrowRequest
	if err := validate.BindAndValidateAny(c, &request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			apperr.NewAppErr(apperr.BAD_REQUEST, fmt.Sprintf("cannot unmarshall request object: %v", err)))
		return
	}

	if err := bc.borrowService.ReturnBorrowedBook(request.BookID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot return borrowed book: %v", err)))
		return
	}

	c.Status(http.StatusOK)
}
