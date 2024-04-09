package service

import (
	"github.com/bastean/bingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bingo/pkg/context/user/domain/model"
)

var IncorrectPassword = errors.InvalidValue{Message: "Password Incorrect"}

func IsPasswordInvalid(hashing model.Hashing, hashed, plain string) {
	if hashing.IsNotEqual(hashed, plain) {
		panic(IncorrectPassword)
	}
}
