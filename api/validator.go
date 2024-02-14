package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/simpleBank/util"
)

var validatorCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// Check currency is supported
		return util.IsSupportCurrency(currency)
	}
	return false
}
