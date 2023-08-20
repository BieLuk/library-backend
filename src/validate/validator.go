package validate

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func Init() {
	Validate = validator.New()
}

func BindAndValidateJson(c *gin.Context, obj any) error {
	if Validate == nil {
		Init()
	}

	if err := c.ShouldBindJSON(&obj); err != nil {
		return fmt.Errorf("error binding object: %w", err)
	}

	if err := Validate.Struct(obj); err != nil {
		return fmt.Errorf("error validating object: %w", err)
	}

	return nil
}
