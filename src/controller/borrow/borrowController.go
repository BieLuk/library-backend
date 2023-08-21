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
// @Summary      Create borrow
// @Description  Creates borrow of given book
// @Tags         Borrows
// @Accept       json
// @Produce      json
// @Param        createBorrowRequest body dto.CreateBorrowRequest true "createBorrowRequest"
// @Success      200  {object}  dto.CreateBorrowResponse
// @Failure      400  {object}  apperr.AppErr
// @Router       /borrows/ [post]
func (bc *borrowController) CreateBorrow(c *gin.Context) {
	var request dto.CreateBorrowRequest
	if err := validate.BindAndValidateJson(c, &request); err != nil {
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
// @Summary      Checks if book with given id is already borrowed
// @Description  Checks if book with given id is already borrowed
// @Tags         Borrows
// @Accept       json
// @Produce      json
// @Param        id   path string  true  "Book ID"
// @Success      200  {object}  dto.IsBookBorrowedResponse
// @Failure      400  {object}  apperr.AppErr
// @Router       /borrows/check/{id} [get]
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

// ReturnBorrowedBook updates brought_date field for borrow with given book id
// @Summary      Updates broughtDate field for borrow with given book id
// @Description  Updates broughtDate field for borrow with given book id
// @Tags         Borrows
// @Accept       json
// @Produce      json
// @Param        returnBorrowRequest body dto.ReturnBorrowRequest true "returnBorrowRequest"
// @Success      200
// @Failure      400  {object}  apperr.AppErr
// @Router       /borrows/return/ [put]
func (bc *borrowController) ReturnBorrowedBook(c *gin.Context) {
	var request dto.ReturnBorrowRequest
	if err := validate.BindAndValidateJson(c, &request); err != nil {
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
