package pretty

import "fmt"
import "github.com/ghodss/yaml"

func YAML(v interface{}) {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}

func Yaml(v interface{}) {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}
