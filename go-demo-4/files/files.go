package files

import (
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, namefile string) {
	file, err := os.Create(namefile)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		panic(err.Error())
	}
}
