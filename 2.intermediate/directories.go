package intermediate

import "os"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func directories() {
	// err := os.Mkdir("subdir", 0755)
	// checkError(err)
	// checkError(os.Mkdir("subdir1", 0755))
	// defer os.RemoveAll("subdir1")

	// os.WriteFile("subdir1/file", []byte(""), 0755)

	// checkError(os.MkdirAll("subdir/parent/child1", 0755))

	// filepath.Walk and filepath.WalkDir

	CheckError(os.RemoveAll("./subdir1"))
}
