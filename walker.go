package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Walker struct {
	Endpoints map[string][]byte
	sourceDir string
}

func NewWalker(sourceDir string) *Walker {
	return &Walker{make(map[string][]byte), sourceDir}
}

func (w *Walker) WalkDirs(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	matched, err := filepath.Match("*.json", info.Name())
	if err != nil {
		return err
	}

	if !matched {
		return nil
	}

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	urlPath := strings.Replace(path, w.sourceDir, "", 1)
	urlPath = strings.Replace(urlPath, info.Name(), "", 1)
	log.Printf("Mapping file %s to path %s", info.Name(), urlPath)
	w.Endpoints[urlPath] = bytes
	return nil
}
