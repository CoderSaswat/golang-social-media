package pkg

import "errors"

var ErrUserExitsWithProvidedEmail error = errors.New("user already exists with the same email")
