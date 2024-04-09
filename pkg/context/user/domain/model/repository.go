package model

import (
	"github.com/bastean/bingo/pkg/context/user/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/user/domain/valueObject"
)

type RepositorySearchFilter struct {
	Id    *valueObject.Id
	Email *valueObject.Email
}

type Repository interface {
	Save(user *aggregate.User)
	Update(user *aggregate.User)
	Delete(id *valueObject.Id)
	Search(filter RepositorySearchFilter) *aggregate.User
}
