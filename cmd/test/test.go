package main

import (
	test "github.com/bastean/bingo/pkg/context/binary/application/create"
)

func main() {
	query := test.NewQuery("", "")
	create := test.NewCreate()
	handler := test.NewQueryHandler(create)

	handler.Handle(query)
}
