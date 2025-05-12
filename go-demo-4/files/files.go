package files

import (
	"os"
)

func ReadFile() {}

func WriteFile(content, namefile string) {
	file, err := os.Create(namefile)
	if err != nil {
		file.Close()
		panic(err.Error())
	}
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		panic(err.Error())
	}
	file.Close()
}
