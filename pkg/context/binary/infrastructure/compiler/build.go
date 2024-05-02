package compiler

import (
	_ "embed"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bastean/bingo/pkg/context/binary/domain/aggregate"
	"github.com/bastean/bingo/pkg/context/binary/domain/builder"
	"github.com/bastean/bingo/pkg/context/binary/domain/service"
	"github.com/bastean/bingo/pkg/context/binary/domain/valueObject"
)

//go:generate rm -rf app/model
//go:generate cp -rp ../../domain/model app

//go:generate touch app/config/config.json

//go:generate bash -c "cd app && go mod tidy && go mod vendor"

//go:generate rm -rf embed
//go:generate mkdir embed
//go:generate bash -c "cd app && tar -czf ../embed/app.tar.gz *"

const tempDirName = "temp"
const tempBuildDirName = "build-*"

const compressedAppFilename = "app.tar.gz"
const entrypointFilename = "main.go"
const configurationFilename = "config.json"

//go:embed embed/app.tar.gz
var compressedAppEmbedded []byte

type BinaryBuild struct {
	TempDir string
}

func (build *BinaryBuild) runCommand(env []string, args ...string) {
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Env = os.Environ()

	cmd.Env = append(cmd.Env, env...)

	if err := cmd.Run(); err != nil {
		service.FailOnError(err, strings.Join(args, " "))
	}
}

func (build *BinaryBuild) createTempDir() {
	os.Mkdir(tempDirName, os.ModePerm)

	tempDir, err := os.MkdirTemp(tempDirName, tempBuildDirName)

	service.FailOnError(err, "failed to create temp directory")

	build.TempDir = tempDir
}

func (build *BinaryBuild) extractEmbeddedApp() {
	compressedAppFile := filepath.Join(build.TempDir, compressedAppFilename)

	err := os.WriteFile(compressedAppFile, compressedAppEmbedded, os.ModePerm)

	service.FailOnError(err, "failed to write embed compressed app")

	build.runCommand(nil, "tar", "-C", build.TempDir, "-xf", compressedAppFile)

	build.runCommand(nil, "rm", "-f", compressedAppFile)
}

func (build *BinaryBuild) writeAppConfigFile(config []byte) {
	configFile := filepath.Join(build.TempDir, "config", configurationFilename)

	err := os.WriteFile(configFile, config, os.ModePerm)

	service.FailOnError(err, "failed to write configuration file")
}

func (build *BinaryBuild) Build(binary *aggregate.Binary) *valueObject.Filepath {
	build.createTempDir()

	build.extractEmbeddedApp()

	binaryConfig, err := json.Marshal(binary)

	service.FailOnError(err, "failed to encoding json")

	build.writeAppConfigFile(binaryConfig)

	targetOS := "GOOS=" + binary.Platform.OS.Value

	targetArch := "GOARCH=" + binary.Platform.Arch.Value

	output := filepath.Join("build", binary.Filename.Value)

	build.runCommand([]string{targetOS, targetArch}, "go", "build", "-C", build.TempDir, "-ldflags", "-s -w", "-mod", "vendor", "-o", output, entrypointFilename)

	filePath := filepath.Join(build.TempDir, output)

	return valueObject.NewFilepath(filePath)
}

func NewBinaryBuilder() builder.Builder {
	return new(BinaryBuild)
}
