package entity

import (
	"fmt"

	"github.com/go-errors/errors"
)

type inner interface {
	error
	ErrorStack() string
}

func innerFromString(template string, parameters ...interface{}) inner {
	return errors.Wrap(fmt.Errorf(template, parameters...), 2)
}
