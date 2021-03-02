package main

import (
	"bytes"
	"fmt"
	"golang.org/x/tools/imports"
	"io/ioutil"
)

func main() {
	// both will print `import "github.com/IlyaFloppy/importsbuglib"`
	// although second one uses github.com/IlyaFloppy/importsbuglib/v2
	stripSingleImportAndProcess("libuserv1/user.go")
	stripSingleImportAndProcess("libuserv2/user.go")

	// and import in usev2.go fixes it
	stripSingleImportAndProcess("libuserv2plus/user.go")

	// aliased import is removed
	stripSingleImportAndProcess("libuseralias/user.go")
}

func stripSingleImportAndProcess(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	linesWithImport := bytes.Split(data, []byte{'\n'})
	lines := make([][]byte, 0, len(linesWithImport))

	for _, line := range linesWithImport {
		if bytes.HasPrefix(line, []byte("import")) {
			continue
		}

		lines = append(lines, line)
	}

	source := bytes.Join(lines, []byte{'\n'}) // does not contain `import "..."`

	processed, err := imports.Process(filename, source, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n-----------------------------------\n", processed)
}
