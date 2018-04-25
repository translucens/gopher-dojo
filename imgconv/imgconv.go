package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/translucens/gopher-dojo/imgconv/dirdigger"
	"github.com/translucens/gopher-dojo/imgconv/imageproc"
	"github.com/translucens/gopher-dojo/imgconv/options"
)

func main() {

	optionError := options.InitOptions()
	if optionError != nil {
		errorHandler(optionError.Error())
	}

	filelist, err := dirdigger.GenerateImageFileListFromDirs(options.Config.InputDirs, options.Config.InputFileType)
	if err != nil {
		errorHandler(err.Error())
	}

	err = imageproc.ConvertImages(filelist, options.Config)
	if err != nil {
		errorHandler(err.Error())
	}

}

func errorHandler(err string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Errorf("%s\n", err)
	os.Exit(1)
}
