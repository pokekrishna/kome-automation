package config

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestValidLoadYaml(t *testing.T) {
	filePath := "/tmp/randomfile123.yaml"

	// Valid YAML File
	if err := createDummyYAMLFile(filePath, "a: b\nc: d"); err != nil {
		t.Errorf("Cannot create dummy file '%s', cannot test\n%v", filePath, err)
	}
	defer deleteFileOrEmptyDir(filePath)

	got, err := loadYAML(filePath)
	if err != nil {
		t.Errorf("'%s' is a valid yaml but test suggests otherwise\n%v",
			filePath, err)
	}
	want := map[interface{}]interface{} {
		"a" : "b",
		"c" : "d",
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Loaded Yaml is not same as expected\nwant: %v\ngot: %v",
			want,got,
		)
	}

}

func TestInvalidLoadYaml(t *testing.T) {
	filePath := "/tmp/randomfile123.yaml"

	// INvalid YAML File
	if err := createDummyYAMLFile(filePath, "a: b\nfoo"); err != nil {
		t.Errorf("Cannot create dummy file '%s', cannot test\n%v", filePath, err)
	}
	defer deleteFileOrEmptyDir(filePath)

	if _, err := loadYAML(filePath); err == nil {
		t.Errorf("'%s' is an INvalid yaml but test suggests otherwise\n%v",
			filePath, err)
	}

}

func createDummyYAMLFile(filePath string, content string) error {
	contentByte := []byte(content)
	err := ioutil.WriteFile(filePath, contentByte, 0644)
	if err != nil {
		return err
	}

	return nil
}

func deleteFileOrEmptyDir(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatalf("Cannot delete file '%s'\n%v", path, err)
	}
}
