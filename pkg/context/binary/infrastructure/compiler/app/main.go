package main

import (
	_ "embed"
	"encoding/json"

	"github.com/bastean/bingo/cmd"
	"github.com/bastean/bingo/model"
)

//go:embed config.json
var Config []byte

func main() {
	root := new(model.Command)

	json.Unmarshal(Config, root)

	cmd.Execute()
}
