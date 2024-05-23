package repositories

import "errors"

var (
	ErrInternalError = errors.New("internal error: ")
	ErrNotCommented  = errors.New("post not commented")
)
