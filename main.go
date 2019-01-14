package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	os.Exit(tree(os.Stdout, os.Args))
}

func tree(w io.Writer, args []string) int {
	// 引数が存在しない場合エラー
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "need to specify target directory")
		return 1
	}
	// 引数のファイルが存在しない場合エラー
	f, err := os.Open(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	// 引数がディレクトリでない場合エラー
	fi, _ := f.Stat()
	if !fi.IsDir() {
		fmt.Fprintln(os.Stderr, "target must be directory")
		return 1
	}

	fmt.Fprintln(w, args[1])
	dirwalk(w, args[1], "")
	return 0
}

func dirwalk(w io.Writer, dir string, bGraph string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		graph := "├── "
		if i+1 == len(files) {
			graph = "└── "
		}
		fmt.Fprintf(w, "%s%s%s\n", bGraph, graph, file.Name())
		if file.IsDir() {
			nGraph := bGraph + "│   "
			if i+1 == len(files) {
				nGraph = bGraph + "    "
			}
			dirwalk(w, filepath.Join(dir, file.Name()), nGraph)
		}
	}
}
