package files

import "os"

type JsonDB struct {
	file string
}

func NewJsonDB(fileName string) *JsonDB {
	return &JsonDB{
		file: fileName,
	}
}

func (j *JsonDB) Read(name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func (j *JsonDB) Write(name string, data []byte) {
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
