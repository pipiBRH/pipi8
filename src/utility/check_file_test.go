package utility

import (
	"testing"
)

func TestCheckMod(t *testing.T) {
	var dirPath = "../flag"
	var filePath = "../flag/flag.go"

	if checkMod(&dirPath) == "dire" {
		t.Log("utility.AcheckModdd PASS")
	} else {
		t.Error("utility.AcheckModdd FAIL")
	}

	if checkMod(&filePath) == "file" {
		t.Log("utility.AcheckModdd PASS")
	} else {
		t.Error("utility.AcheckModdd FAIL")
	}
}
