package main

import (
	"bytes"
	"testing"
)

func TestTree(t *testing.T) {
	output := new(bytes.Buffer)
	expect := `./testdata
├── file
├── test1
│   ├── file1-1
│   ├── file1-2
│   └── test1-1
│       ├── file-1-1-1
│       └── file-1-1-2
└── test2
    ├── file2-1
    └── file2-2
`

	tree(output, []string{"", "./testdata"})
	outputString := output.String()
	if outputString != expect {
		t.Errorf("get:\n%s\nwant:\n%s\n", outputString, expect)
	}
}
