package utility

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"path"
)

func CheckTargetType(targetPath *string) map[string]string {
	var template = make(map[string]string)
	switch checkMod(targetPath) {
	case "dir":
		files, err := ioutil.ReadDir(*targetPath)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			subFilePath := path.Join(*targetPath, f.Name())
			if checkMod(&subFilePath) == "file" {
				template[f.Name()] = subFilePath	
			}
		}
	case "file":
		file := filepath.Base(*targetPath)
		template[file] = *targetPath	
	}
	return template
}

func checkMod(targetPath *string) (result string) {
	fi, err := os.Lstat(*targetPath)
	if err != nil {
		log.Fatal(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		result = "file"
	case mode.IsDir():
		result = "dir"
	default:
		result = ""
		log.Fatal("Not a file or folder")
	}
	return result
}
