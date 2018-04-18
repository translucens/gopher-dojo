package fileproc

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

const separator = filepath.Separator

func TestGenerateJPEGFileList(t *testing.T) {

	actualList, err := GenerateImageFileListFromDirs([]string{"../"}, "jpg")

	fileinfo, err := os.Stat(".")
	if err != nil {
		t.Fatal()
	}
	fmt.Println(fileinfo.Name())

	expected := []string{
		fmt.Sprintf("..%ctestimages%cblack16px.jpg", separator, separator),
		fmt.Sprintf("..%ctestimages%cblack16px-prograssive.jpeg", separator, separator),
		fmt.Sprintf("..%ctestimages%cwhite32px.JPG", separator, separator)}

	sort.Strings(actualList)
	sort.Strings(expected)

	if err != nil {
		t.Fatal()
	}

	if len(actualList) != len(expected) {
		t.Errorf("Length of actual file list %d is not match with expected %d", len(actualList), len(expected))
		t.Log("Actual", actualList)
		t.Log("Expected", expected)
		return
	}

	for i, v := range expected {
		if actualList[i] != v {
			t.Errorf("Filepath is not match; actual %s; expected %s", actualList[i], v)
		}
	}

}

func TestGeneratePNGFileList(t *testing.T) {

	actualList, err := GenerateImageFileListFromDirs([]string{"../"}, "png")
	expected := []string{
		fmt.Sprintf("..%ctestimages%cblack16px.png", separator, separator),
		fmt.Sprintf("..%ctestimages%cchilddir%cwhite32px.png", separator, separator, separator)}

	sort.Strings(actualList)
	sort.Strings(expected)

	if err != nil {
		t.Fatal()
	}

	if len(actualList) != len(expected) {
		t.Errorf("Length of actual file list %d is not match with expected %d", len(actualList), len(expected))
		t.Log("Actual", actualList)
		t.Log("Expected", expected)
		return
	}

	for i, v := range expected {
		if actualList[i] != v {
			t.Errorf("Filepath is not match; actual %s; expected %s", actualList[i], v)
		}
	}

}
