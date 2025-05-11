package commonsValidator

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"time"
)

func unixTimestampValidator(fl validator.FieldLevel) bool {
	timestamp := fl.Field().Int()
	now := time.Now().Unix()
	m := now + 10*365*24*60*60
	return timestamp > 0 && timestamp < m
}

func RegisterCustomValidators(r *gin.Engine) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("unix_ts", unixTimestampValidator)
		if err != nil {
			return
		}
	}
}
