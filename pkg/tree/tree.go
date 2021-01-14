package tree

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YamlFile struct {
	fields map[string]interface{}
}

func (y *YamlFile) Print() string {
	return "this is a file"
}

// Generate creates the tree of a structured file
func Generate(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	yamlFile := YamlFile{}
	err = yaml.Unmarshal(file, &yamlFile)
	if err != nil {
		return err
	}

	fmt.Println(yamlFile.Print())
	return nil
}
