package main

import "errors"

var (
	ErrRequestBodyNotProperlyFormatted = errors.New("Request body not properly formatted")
	ErrInternalServerError             = errors.New("Something went wrong")
)
