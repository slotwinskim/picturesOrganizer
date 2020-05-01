package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	root := "./"
	tree(root)

}

func tree(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("Error")
		return
	}

	for nr, file := range files {
		if !file.IsDir() {
			pre := "├─── "
			if nr == len(files)-1 {
				pre = "└─── "
			}

			fullName := pre + getProperName(file.Name())
			fmt.Println(fullName)
		} else {
			fmt.Println(getProperName(file.Name()))
			tree(path + file.Name())
		}
	}
}

func getProperName(name string) string {
	if !strings.Contains(name, " ") {
		return name
	} else {
		return "\"" + name + "\""
	}
}
