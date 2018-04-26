package dirdigger_test

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/translucens/gopher-dojo/imgconv/dirdigger"
)

const Separator = string(filepath.Separator)
const TestdataDir = ".." + Separator + "testdata" + Separator

func TestGenerateJPEGFileList(t *testing.T) {

	actualList, err := dirdigger.GenerateImageFileListFromDirs([]string{"../"}, "jpg")

	fileinfo, err := os.Stat(".")
	if err != nil {
		t.Fatal()
	}
	fmt.Println(fileinfo.Name())

	expected := []string{
		fmt.Sprintf(TestdataDir + "black16x32px.jpg"),
		fmt.Sprintf(TestdataDir + "black16px-prograssive.jpeg"),
		fmt.Sprintf(TestdataDir + "white32px.JPG"),
		fmt.Sprintf(TestdataDir + "invalid.jpg")}

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

	actualList, err := dirdigger.GenerateImageFileListFromDirs([]string{"../"}, "png")
	expected := []string{
		fmt.Sprintf(TestdataDir + "black16px.png"),
		fmt.Sprintf(TestdataDir + "alphach.png"),
		fmt.Sprintf(TestdataDir+"childdir%swhite32px.png", Separator)}

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

func TestGenerateUnsupportedFileList(t *testing.T) {
	if _, err := dirdigger.GenerateImageFileListFromDirs([]string{"../"}, "gif"); err == nil {
		t.Fatalf("GIF should be not supported in this program!")
	}
}
