package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	strcase "github.com/iancoleman/strcase"
)

func replaceContentCases(content string, name string) (result string) {
	result = strings.ReplaceAll(content, "{{name}}", name)
	result = strings.ReplaceAll(result, "{{nameCamel}}", strcase.ToLowerCamel(name))
	result = strings.ReplaceAll(result, "{{NameCamel}}", strcase.ToCamel(name))
	result = strings.ReplaceAll(result, "{{nameKebab}}", strcase.ToKebab(name))
	result = strings.ReplaceAll(result, "{{NameKebab}}", strcase.ToScreamingKebab(name))
	result = strings.ReplaceAll(result, "{{Name}}", strings.Title(name))
	result = strings.ReplaceAll(result, "{{NAME}}", strings.ToUpper(name))
	return result
}

func CreateFiles(config CreateFileConfig) (results []CreateFileResult) {
	var files []string

	err := filepath.Walk(config.baseTemplatePath, func(path string, info os.FileInfo, err error) error {
		if path != config.baseTemplatePath && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup

	for _, filePath := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			input, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalln(err)
			}

			content := string(input)
			results = append(
				results, CreateFile(
					replaceContentCases(strings.Replace(file, config.baseTemplatePath, config.basePath, 1), config.name),
					replaceContentCases(content, config.name),
				))

		}(filePath)

	}

	wg.Wait()

	return results
}
