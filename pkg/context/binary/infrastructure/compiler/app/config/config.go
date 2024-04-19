package config

import (
	_ "embed"
	"encoding/json"

	"github.com/bastean/bingo/model"
)

//go:embed config.json
var configEmbed []byte

var Root = new(model.Command)

func init() {
	json.Unmarshal(configEmbed, Root)
}
