package cmd

import (
	"flag"
	"log"
	"path/filepath"
)

var ImageTag *string
var FilePath *string

func init() {
	ImageTag = flag.String("t", "latest", "specific image tag")
	filePath := flag.String("f", "k8s", "specific yaml folder")
	flag.Parse()
	tmpPath, err := filepath.Abs(*filePath)
	FilePath = &tmpPath
	if err != nil {
		log.Fatal(err)
	}
}
