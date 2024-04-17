package compiler

import (
	_ "embed"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/builder"
	"github.com/bastean/bingo/pkg/context/binary/domain/service"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

//go:generate rsync -a --delete ../../domain/model app
//go:generate touch app/config.json
//go:generate bash -c "cd app && go mod tidy && go mod vendor"
//go:generate rm -rf embed
//go:generate mkdir embed
//go:generate bash -c "cd app && tar -czf ../embed/app.tar.gz *"

const dirTempName = "temp"
const dirTempBuildName = "build-*"

const fileCompressAppName = "app.tar.gz"
const fileEntrypointName = "main.go"
const fileConfigurationName = "config.json"

//go:embed embed/app.tar.gz
var compressedApp []byte

type Binary struct{}

func (binary *Binary) Build(root *aggregate.Root) *valueObject.FilePath {
	os.Mkdir(dirTempName, os.ModePerm)

	dirTemp, err := os.MkdirTemp(dirTempName, dirTempBuildName)

	service.FailOnError(err, "failed to create temp directory")

	compressedAppFile := filepath.Join(dirTemp, fileCompressAppName)

	err = os.WriteFile(compressedAppFile, compressedApp, os.ModePerm)

	service.FailOnError(err, "failed to write embed compressed app")

	service.RunCommand("tar", "-C", dirTemp, "-xf", compressedAppFile)

	service.RunCommand("rm", "-f", compressedAppFile)

	rootJson, err := json.Marshal(root)

	service.FailOnError(err, "failed to encoding json")

	configFile := filepath.Join(dirTemp, fileConfigurationName)

	err = os.WriteFile(configFile, rootJson, os.ModePerm)

	service.FailOnError(err, "failed to write configuration file")

	output := filepath.Join("build", root.Command.Name)

	service.RunCommand("go", "build", "-C", dirTemp, "-mod", "vendor", "-o", output, fileEntrypointName)

	filePath := filepath.Join(dirTemp, output)

	return valueObject.NewFilePath(filePath)
}

func NewBinaryBuilder() builder.Builder {
	return new(Binary)
}