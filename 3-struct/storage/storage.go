package storage

type Storage struct {
	files []string
}

func NewStorage() *Storage {
	return &Storage{
		files: []string{},
	}
}
