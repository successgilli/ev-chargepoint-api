package chargepoint

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// RegisterCustomValidators registers custom validators 
func RegisterCustomValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("unable to find validator engine")
	}

	if err := v.RegisterValidation("chargepoint_status", ChargePointStatusValidator); err != nil {
		return err
	}

	return nil
}
