package valueObjectMother

import (
	"github.com/bastean/bingo/pkg/context/binary/domain/service/mother"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

func RandomFilepath() *valueObject.Filepath {
	return valueObject.NewFilepath(mother.Create.Regex(`\/{0,1}[\w-.()]{1,10}\/{1}[\w-.()]{1,10}`))
}

func InvalidFilepath() *valueObject.Filepath {
	return valueObject.NewFilepath("//")
}

func EmptyFilepath() *valueObject.Filepath {
	return valueObject.NewFilepath("")
}
