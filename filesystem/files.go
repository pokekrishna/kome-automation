package filesystem

import (
	"fmt"
	"os"
)

// FileExists tells if file exists and has permissions or not?
func FileExists(filePath string) error {
	fileInfo, err := os.Stat(filePath)

	//valid file path?
	if os.IsNotExist(err) {
		return err
	}

	// is permission?
	// TODO: The following function does not tell if the file is readable or not
	// example, even if the permission is 0000, the following function will not error
	if os.IsPermission(err) {
		return err
	}

	// is file? (and not dir!)
	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", filePath)
	}

	return nil
}
