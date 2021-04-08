package service

import (
	"github.com/pkg/errors"
)

var (
	ErrProductNotFound = errors.New("Product Not Found")
	ErrProductInvalid  = errors.New("Product Invalid")
)
