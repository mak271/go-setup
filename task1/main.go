package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func filesLoop(root string, files []os.DirEntry, tree *string, printFiles bool, prefix string) error {
	slices.SortFunc(files, func(first, second os.DirEntry) int {
		return strings.Compare(first.Name(), second.Name())
	})

	if !printFiles {
		files = slices.DeleteFunc(files, func(file os.DirEntry) bool {
			return !file.IsDir()
		})
	}

	for _, file := range files {

		prefLines := ""

		if files[len(files)-1] == file {
			prefLines = "└───"
		} else {
			prefLines = "├───"
		}

		fileInfo, fileInfoError := file.Info()
		if fileInfoError != nil {
			return fileInfoError
		}

		fileSize := ""

		if printFiles {
			if fileInfo.Size() == 0 {
				fileSize = " (empty)"
			} else {
				fileSize = fmt.Sprint(" (", strconv.Itoa(int(fileInfo.Size())), "b)")
			}
		}

		*tree = *tree + prefix + prefLines + file.Name() + fileSize + "\n"

		if file.IsDir() {
			openRes, openError := os.Open(filepath.Join(root, file.Name()))
			if openError != nil {
				return openError
			}
			openedFiles, oFilesErrror := openRes.ReadDir(0)
			if oFilesErrror != nil {
				return oFilesErrror
			}
			if len(openedFiles) > 0 {

				if files[len(files)-1] == file {
					prefix = prefix + "\t"
				} else {
					prefix = prefix + "│\t"
				}

				loopError := filesLoop(filepath.Join(root, file.Name()), openedFiles, tree, printFiles, prefix)

				prefix, _ = strings.CutSuffix(prefix, "│\t")

				if loopError != nil {
					return loopError
				}
			}
		}
	}

	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	var tree = ""

	files, filesError := os.ReadDir("./testdata")

	if filesError != nil {
		return filesError
	}

	loopError := filesLoop(path, files, &tree, printFiles, "")

	if loopError != nil {
		return loopError
	}

	fmt.Fprintln(out, tree)
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := "testdata"
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
