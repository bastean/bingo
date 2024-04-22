package builder

import (
	"github.com/bastean/bingo/pkg/context/binary/application/create"
	"github.com/bastean/bingo/pkg/context/binary/infrastructure/compiler"
)

var binaryBuilder = compiler.NewBinaryBuilder()
var binaryCreate = create.NewCreate(binaryBuilder)
var BinaryCreateHandler = create.NewQueryHandler(binaryCreate)
