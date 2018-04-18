package main

import (
	"os"

	"github.com/translucens/gopher-dojo/imgconv/fileproc"
	"github.com/translucens/gopher-dojo/imgconv/options"
)

func main() {

	optionError := options.InitOptions()
	if optionError != nil {
		errorHandler(optionError.Error())
	}

	filelist, err := fileproc.GenerateImageFileListFromDirs(options.Config.InputDirs, options.Config.InputFileType)
	if err != nil {
		errorHandler(err.Error())
	}

	err = fileproc.ConvertImages(filelist, options.Config)
	if err != nil {
		errorHandler(err.Error())
	}

}

func errorHandler(err string) {
	os.Stderr.WriteString(err)
	os.Stderr.WriteString("\n")
	os.Exit(1)
}
