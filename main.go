package main

import (
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
	"os"
	"path"
	"time"
)

type CreateFileResult struct {
	path  string
	error error
}

type CreateFileConfig struct {
	name             string
	basePath         string
	baseTemplatePath string
}

type Flags struct {
	noDir         bool
	templatePath  string
	childTemplate string
}

func isInNodeModules(current string) bool {
	if len(current) >= 17 && current[len(current)-17:] == "node_modules/.bin" {
		return true
	}
	return false
}

func currentExecutablePath() (string, error)  {
	ex, err := os.Executable()
	if err != nil { return "", err}
	return path.Dir(ex), nil
}

func main() {

	var err error = nil

	flags := new(Flags)

	args := os.Args[1:]

	flag.BoolVar(&flags.noDir, "no-dir", false, "no create directory")

	flag.StringVar(&flags.templatePath, "base-template", "", "manually set template path")

	flag.StringVar(&flags.childTemplate, "t", "", "set child template dir")

	flag.Parse()

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	var templatePath = path.Join(flags.templatePath, flags.childTemplate)

	currentPath, err := currentExecutablePath()

	if err != nil {
		log.Fatal(err)
	}

	if isInNodeModules(currentPath) && flags.templatePath == "" {
		templatePath = path.Join(fmt.Sprintf("%s%s", currentPath[:len(currentPath)-17], "template"), flags.childTemplate)
	}

	if templatePath == "" {
		templatePath = "template"
	}

	if len(os.Args) <= 1 {
		log.Fatal(aurora.Red("Require dir name..."))
	}

	basePath := dir

	start := time.Now()

	for _, arg := range args {
		if arg[:1] != "-" {
			if !flags.noDir {
				err = CreateDir(path.Join(arg))
				if err != nil {
					break
				}
				basePath = path.Join(dir, arg)
			}

			fmt.Println(aurora.Magenta("Base path: "), aurora.Cyan(basePath))
			fmt.Println(aurora.Magenta("Template path: "), aurora.Cyan(templatePath))

			config := CreateFileConfig{
				name:             arg,
				basePath:         basePath,
				baseTemplatePath: templatePath,
			}

			results := CreateFiles(config)

			fmt.Println(aurora.Magenta("Result:"))

			for _, result := range results {
				if result.error == nil {
					fmt.Println(aurora.Green("- Created file "), aurora.Cyan(result.path))
				} else {
					fmt.Println(aurora.Red("- Failed create file "), aurora.Cyan(result.path))
					fmt.Println(fmt.Sprintf("-- %s", result.error.Error()))
				}
			}
		}

	}

	fmt.Println(aurora.Cyan("Finish"), aurora.Green(fmt.Sprintf("%fms", float64(time.Since(start)) / float64(time.Millisecond))))

	if err != nil {
		log.Fatal(aurora.Red(err))
	}
}
