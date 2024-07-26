package api

import (
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Input struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	RedirectURI string `json:"redirect_uri" binding:"url"`
	Username string `json:"user_name" binding:"username_validate"`
}

var usernameValidate validator.Func = func(fl validator.FieldLevel) bool {
	if v, ok := fl.Field().Interface().(string); ok {
		if v == "anonymous" {
			return false
		}
	}
	return true
}

func GetService(number string) string {
	china := []string {"+86", "+0086"}
	for _, prefix := range china {
		if strings.HasPrefix(number, prefix) {
			return "YiroamingService"
		}
	}
	return "CyoloService"
}

func SmsGateway(c *gin.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("username_validate", usernameValidate)
	}
	var input Input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	c.JSON(http.StatusOK, gin.H{"service": GetService(input.PhoneNumber)})
	// c.Writer.WriteHeader(http.StatusNoContent)
}