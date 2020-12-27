package dbmlgorm

import (
	"fmt"
	"github.com/duythinht/dbml-go/core"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
)

type Config struct {
	Package string
	Imports []string
	Types   []string
}

type Type struct {
	original string
	new      string
}

func getProjectConfig(dbml *core.DBML) (string, map[string]string) {
	var config Config
	yamlBytes := []byte(dbml.Project.Note)
	err := yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	str := fmt.Sprintf(""+
		"// Auto-generated by dbml-to-gorm\n"+
		"// Do not edit by hand!\n"+
		"// Project: %v\n\n"+
		"package %v\n\n",
		dbml.Project.Name, config.Package)
	if len(config.Imports) > 0 {
		str += "import(\n"
		for _, imp := range config.Imports {
			str += "\t\"" + imp + "\"\n"
		}
		str += ")\n\n"
	}

	types := make(map[string]string)
	for _, t := range config.Types {
		splited := strings.Split(t, ":")
		types[splited[0]] = splited[1]
	}
	return str, types
}
