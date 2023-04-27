package internalerrors

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternal error = errors.New("internal error")

func ProcessErrorToReturn(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternal
	}
	return err
}
