package handlers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func createValidationError(err error) error {
	// Convert to ValidationErrors type
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	// Build array with incorrect error fields and create error string
	var incorrectFields []string
	for _, e := range validationErrors {
		incorrectFields = append(incorrectFields, e.Field()) // TODO: Return JSON fields instead of Go naming
	}
	incorrectFieldsString := strings.Join(incorrectFields, ", ")
	formattedError := fmt.Errorf("incorrect or missing fields in body: %s", incorrectFieldsString)

	// Return formatted error string
	return formattedError
}

func (h Handlers) extractBody(c *gin.Context, data interface{}) {
	// Bind the request body
	if err := binding.JSON.Bind(c.Request, data); err != nil {
		h.sendInvalidPostBody(c, err)
		c.Abort()
		return
	}

	// Validate the body with `validate` struct tags
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		// When error is found, create user friendly error with the incorrect fields and send 400 response
		h.sendInvalidPostBody(c, createValidationError(err))
		c.Abort()
		//h.Logger.Debug(validationErr.Error())
		return
	}
}

func (h Handlers) checkResponseBody(c *gin.Context, data interface{}) {
	if data == nil {
		return // When data is nil, do not run validation
	}

	// Validate the response body struct
	validate := validator.New()
	rt := reflect.TypeOf(data)
	var err error
	switch rt.Kind() {
	case reflect.Slice:
		err = validate.Var(data, "dive") // TODO: Test better
	default:
		err = validate.Struct(data)
	}

	// When error is found, create user friendly error with the incorrect fields and send 500 response
	if err != nil {
		h.sendServerError(c, createValidationError(err))
		c.Abort()
	}
}
