package netsdk

import "errors"

var (
	ErrLoginFailed    = errors.New("login failed")
	ErrInvalidAddress = errors.New("invalid address format, must have ip and port")
)
