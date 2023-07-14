package pantrygo

import "errors"

var errx = struct {
	ErrPantry           error
	ErrUnknown          error
	ErrResourceNotFound error
}{
	ErrPantry:           errors.New("errorPantry"),
	ErrUnknown:          errors.New("errorUnknown"),
	ErrResourceNotFound: errors.New("resource_not_found"),
}
