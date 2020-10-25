package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/logrusorgru/aurora"
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
	dir           string
	templatePath  string
	childTemplate string
}

func isInNodeModules(current string) bool {
	if len(current) >= 17 && current[len(current)-17:] == "node_modules/.bin" {
		return true
	}
	return false
}

func currentExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return path.Dir(ex), nil
}

func main() {

	var err error = nil

	flags := new(Flags)

	args := os.Args[1:]

	if err != nil {
		log.Fatal(err)
	}

	flag.BoolVar(&flags.noDir, "no-dir", false, "no create directory")

	flag.StringVar(&flags.dir, "dir", "", "create template in path")

	flag.StringVar(&flags.templatePath, "base-template", path.Join("template"), "manually set template path location")

	flag.StringVar(&flags.childTemplate, "t", "", "set child template dir")

	flag.Parse()

	basePath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	dir := path.Join(basePath, flags.dir)

	var templatePath = path.Join(flags.templatePath, flags.childTemplate)

	if len(os.Args) <= 1 {
		log.Fatal(aurora.Red("Require dir name..."))
	}

	start := time.Now()

	for _, arg := range args {
		if arg[:1] != "-" {
			if !flags.noDir {
				err = CreateDir(path.Join(dir, arg))
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

	fmt.Println(aurora.Cyan("Finish"), aurora.Green(fmt.Sprintf("%fms", float64(time.Since(start))/float64(time.Millisecond))))

	if err != nil {
		log.Fatal(aurora.Red(err))
	}
}
