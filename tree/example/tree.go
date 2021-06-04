package example

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const LineBreak = "\n"
var ignoreFirstPath = "testdata"

func DirTree(buf *bytes.Buffer, path string, indent string, displaySize bool) error{
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatal(fi)
	}

	checkDisplaySize(displaySize, fi, buf)

	if !fi.IsDir() {
		return nil
	}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	var filesName []string
	for _, file := range files {
		switch  {
		case displaySize:
			filesName = append(filesName,file.Name())
		case !displaySize && file.IsDir():
			filesName = append(filesName,file.Name())
		}
	}

	for i, name := range filesName {
		add := "│	"
		if i == len(filesName) - 1 {
			buf.WriteString(indent + "└───")
			add = "	"
		} else {
			buf.WriteString(indent + "├───")
		}
		if err:= DirTree(buf, filepath.Join(path,name), concatStrings(indent,add),displaySize); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}


func concatStrings(str1, str2 string) string{
	var newStr strings.Builder

	newStr.WriteString(str1)
	newStr.WriteString(str2)

	return newStr.String()
}

func checkDisplaySize(displaySize bool, fi fs.FileInfo, buf *bytes.Buffer)  {
	if displaySize {
		if fi.Name() != ignoreFirstPath {
			if fi.IsDir() {
				buf.WriteString(fi.Name() + LineBreak)
			} else {
				if (fi.Size() == 0) {
					buf.WriteString(fi.Name() + " (empty)" + LineBreak)
				} else {
					buf.WriteString(fi.Name() + " (" + strconv.FormatInt(fi.Size(), 10) + "b)" + LineBreak)
				}
			}
		}
	} else {
		if fi.Name() != ignoreFirstPath {
			if fi.IsDir() {
				buf.WriteString(fi.Name() + LineBreak)
			}
		}
	}
}
