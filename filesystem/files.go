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
	if os.IsPermission(err){
		return err
	}

	// is file? (and not dir!)
	if !fileInfo.Mode().IsRegular(){
		return fmt.Errorf("%s is not a regular file.", filePath)
	}

	return nil
}


