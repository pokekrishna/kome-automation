package filesystem

import (
	"log"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	path := "./randomfilename123"
	err := FileExists(path)
	if err == nil {
		t.Errorf("File '%s' should not exist but it does.\n %v", path, err)
	}

	err = os.Mkdir(path, 0750)
	if err != nil {
		log.Fatalf("Cannot create dummy dir '%s'\n%v", path, err)
	}
	defer deleteFileOrEmptyDir(path)

	if err := FileExists(path); err == nil {
		t.Errorf("Path '%s' is a dir., FileExists() should throw error.", path)
	}

}

func deleteFileOrEmptyDir(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatalf("Cannot delete file '%s'\n%v", path, err)
	}
}
