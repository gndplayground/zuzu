package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

const acbTsx = `import React, {FunctionComponent} from 'react';

const Acb: FunctionComponent = () => {
    return (
        <h1>Hello!</h1>
    );
};

export default Acb;`

const indexTsx = `export {default} from './Acb'`

const aCBTsx = `acbAcbACB`

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkGeneratedFiles(list []string) (file string, result bool) {
	input, err := ioutil.ReadFile(list[0])

	fatalErr(err)

	if string(input) != acbTsx {
		return list[0], false
	}

	input, err = ioutil.ReadFile(list[1])

	if string(input) != indexTsx {
		return list[1], false
	}

	input, err = ioutil.ReadFile(list[2])

	if string(input) != aCBTsx {
		return list[2], false
	}

	fatalErr(err)

	return "", true
}

func TestZuZu(t *testing.T) {
	t.Run("Should return error 'Require dir name...'", func(t *testing.T) {
		cmd := exec.Command("./zuzu")
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		cmd.Run()
		fmt.Println("out:", outb.String())
		fmt.Println("err:", errb.String())
		if !strings.Contains(errb.String(), "Require dir name...") {
			t.Errorf("Should return error 'Require dir name...'")
		}
	})

	t.Run("Should create template when stand on .bin node", func(t *testing.T) {
		cmd := exec.Command("node_modules/.bin/zuzu", "acb")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		fatalErr(err)

		var files []string

		walks := []string{"acb", "acb/acb.tsx", "acb/index.tsx", "acb/m", "acb/m/ACB.tsx"}

		err = filepath.Walk("acb", func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})

		fatalErr(err)

		sort.Strings(walks)

		sort.Strings(files)
		fmt.Println(walks, files)
		if !reflect.DeepEqual(walks, files) {
			t.Errorf("Should create correct files")
		}

		if fileErr, result := checkGeneratedFiles([]string{"acb/acb.tsx", "acb/index.tsx", "acb/m/ACB.tsx"}); !result {
			t.Errorf("Content generated mismatch in file %s", fileErr)
		}

		err = os.RemoveAll("acb")

		fatalErr(err)
	})

	t.Run("Should create template when stand on .bin node with no dir", func(t *testing.T) {
		cmd := exec.Command("node_modules/.bin/zuzu", "-no-dir", "acb")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		fatalErr(err)

		if fileErr, result := checkGeneratedFiles([]string{"acb.tsx", "index.tsx", "m/ACB.tsx"}); !result {
			t.Errorf("Content generated mismatch in file %s", fileErr)
		}

		err = os.Remove("acb.tsx")

		fatalErr(err)

		err = os.Remove("index.tsx")

		fatalErr(err)

		err = os.RemoveAll("m")

		fatalErr(err)
	})

	t.Run("Should create template when stand on .bin node with specific template", func(t *testing.T) {
		cmd := exec.Command("node_modules/.bin/zuzu", "-t=m", "acb")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		fatalErr(err)

		input, err := ioutil.ReadFile("acb/ACB.tsx")

		if string(input) != aCBTsx {
			t.Errorf("Content generated mismatch in file %s", "acb/ACB.tsx")
		}

		err = os.RemoveAll("acb")

		fatalErr(err)
	})

	t.Run("Should create template when stand on .bin node with custom base template", func(t *testing.T) {
		cmd := exec.Command("node_modules/.bin/zuzu", "-base-template=template2", "acb")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		fatalErr(err)

		input, err := ioutil.ReadFile("acb/acb.tsx")

		if string(input) != "Acb" {
			t.Errorf("Content generated mismatch in file %s", "acb/acb.tsx")
		}

		err = os.RemoveAll("acb")

		fatalErr(err)
	})

	t.Run("Should work when call exec directly", func(t *testing.T) {
		cmd := exec.Command("./zuzu", "acb")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		fatalErr(err)

		var files []string

		walks := []string{"acb", "acb/acb.tsx", "acb/index.tsx", "acb/m", "acb/m/ACB.tsx"}

		err = filepath.Walk("acb", func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})

		fatalErr(err)

		sort.Strings(walks)

		sort.Strings(files)

		if !reflect.DeepEqual(walks, files) {
			t.Errorf("Should create correct files")
		}

		if fileErr, result := checkGeneratedFiles([]string{"acb/acb.tsx", "acb/index.tsx", "acb/m/ACB.tsx"}); !result {
			t.Errorf("Content generated mismatch in file %s", fileErr)
		}

		err = os.RemoveAll("acb")

		fatalErr(err)
	})

	t.Run("Should create template with correct string case", func(t *testing.T) {
		cmd := exec.Command("node_modules/.bin/zuzu", "-base-template=template3", "-t=stringcase", "ContentCase")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		fatalErr(err)

		input, err := ioutil.ReadFile("ContentCase/camel.txt")

		if string(input) != "contentCaseContentCase" {
			t.Errorf("Content generated mismatch in file %s", "ContentCase/camel.txt")
		}

		input, err = ioutil.ReadFile("ContentCase/kebab.txt")

		if string(input) != "content-caseCONTENT-CASE" {
			t.Errorf("Content generated mismatch in file %s", "ContentCase/kebab.txt")
		}

		err = os.RemoveAll("ContentCase")

		fatalErr(err)
	})
}
