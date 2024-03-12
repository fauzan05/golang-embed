package golangembed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestGetEmbed(t *testing.T) {
	fmt.Println(version)
}


//go:embed Livewire.png
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("logo-new.png", logo, fs.ModePerm) //membuat file dengan nama logo-new.png dan isinya sama persis dengan file yang diembed yaitu Livewire.png
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var allFiles embed.FS

func TestPathMatcher(t *testing.T) {
	list, _ := allFiles.ReadDir("files")
	for _, file := range list {
		if !file.IsDir() {
			content, _ := allFiles.ReadFile("files/" + file.Name())
			fmt.Println("Nama file-nya : ", file.Name(), "| Isinya : ", string(content))
		}
	}
}