package domain

import "errors"

var ErrAuthFailed = errors.New("auth failed")
var ErrEmailTaken = errors.New("email is already taken")
