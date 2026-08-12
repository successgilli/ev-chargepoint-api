package chargepoint

import (
	"github.com/go-playground/validator/v10"
	entity "github.com/successgilli/api-app/internal/entities/chargepoint"
)

// ChargePointStatusValidator charge point status validator
var ChargePointStatusValidator validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(entity.ChargePointStatus)
	if !ok {
		return false
	}

	return value.IsValid()
}
