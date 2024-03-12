package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed Livewire.png
var logo []byte

//go:embed files/*.txt
var allFiles embed.FS

func main(){
	fmt.Println(version)
	err := ioutil.WriteFile("logo-new.png", logo, fs.ModePerm) //membuat file dengan nama logo-new.png dan isinya sama persis dengan file yang diembed yaitu Livewire.png
	if err != nil {
		panic(err)
	}
	list, _ := allFiles.ReadDir("files")
	for _, file := range list {
		if !file.IsDir() {
			content, _ := allFiles.ReadFile("files/" + file.Name())
			fmt.Println("Nama file-nya : ", file.Name(), "| Isinya : ", string(content))
		}
	}
}