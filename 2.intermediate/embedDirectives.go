package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed tempfolder
var tempFolder embed.FS

func embedDirectorss() {
	fmt.Println(content)
	cnt, err := tempFolder.ReadFile("tempfolder/example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
	err = fs.WalkDir(tempFolder, "tempfolder", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
