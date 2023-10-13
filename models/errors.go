package models

import "errors"

var (
	ErrPhoneNumberAlreadyRegistered = errors.New("phone number already registered")
	ErrPhoneNumberNotRegistered     = errors.New("user not registered")
	ErrRestaurantAlreadyRegistered  = errors.New("restaurant already registered")
	ErrRestaurantNotRegistered      = errors.New("restaurant not registered")
	ErrUnauthorizedUser             = errors.New("user unauthorized to perform action")
	ErrQuantityLessThanZero         = errors.New("quantity less than zero")
	ErrNotEnoughQuantity            = errors.New("quantity not enough")
	ErrDoesNotDeliver               = errors.New("restaurant does not deliver at location")
)
