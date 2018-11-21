package main

import (
	"mb8/src/flag"
	"mb8/src/template_execute"
	"mb8/src/utility"
)

func main() {
	templateInfo := utility.CheckTargetType(cmd.FilePath)
	execute.ExecuteTemplate(templateInfo, cmd.ImageTag)
}
