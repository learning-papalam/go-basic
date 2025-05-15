package files

import "os"

func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func WriteFile(name string, data []byte) {
	file, err := os.Create(name)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		panic(err.Error())
	}
}
