package options_test

import (
	"testing"

	"github.com/translucens/gopher-dojo/imgconv/options"
)

func TestValidOption(t *testing.T) {

	options := options.Options{

		Overwrite:      true,
		InputFileType:  "jpg",
		OutputFileType: "png",
		InputDirs:      []string{"./"},
		JPEGQuality:    80,
		PNGLevel:       0,
	}

	if err := options.Verify(); err != nil {
		t.Fatalf("Valid options should be passed; %s", err.Error())
	}
}

func TestEmptyInputDir(t *testing.T) {
	if err := options.InitOptions(); err == nil {
		t.Fatalf("Input dir is required but passed check.")
	}
}

func TestDefaultValues(t *testing.T) {
	options.Config.InputDirs = []string{"../"}
	if err := options.Config.Verify(); err != nil {
		t.Fatalf("Default values are should be passed validation.")
	}

}

func TestInvalidOption(t *testing.T) {

	options := options.Options{

		Overwrite:      true,
		InputFileType:  "jpg",
		OutputFileType: "jpeg",
		InputDirs:      []string{"./"},
		JPEGQuality:    101, // invalid
		PNGLevel:       0,
	}

	if err := options.Verify(); err == nil {
		t.Fatalf("Invalid options should not be passed")
	}
}
