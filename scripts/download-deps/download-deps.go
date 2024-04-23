package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Dependency struct {
	Directory   string
	Filename    string
	DownloadURL string
}

const basePath = "pkg/cmd/server/static/vendor"

var Dependencies = []*Dependency{
	{Directory: "htmx.org", DownloadURL: "https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js"},
	{Directory: "htmx.org/ext", DownloadURL: "https://unpkg.com/htmx.org@1.9.12/dist/ext/json-enc.js"},
	{Directory: "htmx.org/ext", DownloadURL: "https://unpkg.com/htmx.org@1.9.12/dist/ext/response-targets.js"},
	{Directory: "htmx.org/ext", DownloadURL: "https://unpkg.com/htmx.org@1.9.12/dist/ext/multi-swap.js"},
	{Directory: "htmx.org/ext", DownloadURL: "https://unpkg.com/htmx.org@1.9.12/dist/ext/remove-me.js"},
	{Directory: "htmx.org/ext", DownloadURL: "https://unpkg.com/htmx.org@1.9.12/dist/ext/alpine-morph.js"},

	{Directory: "alpinejs.dev", Filename: "alpine.min.js", DownloadURL: "https://cdn.jsdelivr.net/npm/alpinejs@3.13.9/dist/cdn.min.js"},
	{Directory: "alpinejs.dev/ext", Filename: "persist.min.js", DownloadURL: "https://cdn.jsdelivr.net/npm/@alpinejs/persist@3.13.9/dist/cdn.min.js"},
	{Directory: "alpinejs.dev/ext", Filename: "morph.min.js", DownloadURL: "https://cdn.jsdelivr.net/npm/@alpinejs/morph@3.13.9/dist/cdn.min.js"},

	{Directory: "jquery.com", DownloadURL: "https://cdn.jsdelivr.net/npm/jquery@3.7.1/dist/jquery.min.js"},

	{Directory: "fomantic-ui.com", DownloadURL: "https://cdn.jsdelivr.net/npm/fomantic-ui@2.9.3/dist/semantic.min.js"},
	{Directory: "fomantic-ui.com", DownloadURL: "https://cdn.jsdelivr.net/npm/fomantic-ui@2.9.3/dist/semantic.min.css"},
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	os.RemoveAll(basePath)

	os.Mkdir(basePath, os.ModePerm)

	for _, dependency := range Dependencies {
		if dependency.Filename == "" {
			dependency.Filename = filepath.Base(dependency.DownloadURL)
		}

		directory := filepath.Join(basePath, dependency.Directory)

		os.Mkdir(directory, os.ModePerm)

		log.Printf("Downloading: %s", dependency.Filename)

		file, err := os.Create(filepath.Join(directory, dependency.Filename))
		failOnError(err, "failed to create file "+dependency.Filename)

		response, err := http.Get(dependency.DownloadURL)
		failOnError(err, "failed to download url "+dependency.DownloadURL)

		_, err = io.Copy(file, response.Body)
		failOnError(err, "failed to copy on file "+dependency.Filename)

		log.Printf("Added: %s", dependency.Filename)
	}
}
