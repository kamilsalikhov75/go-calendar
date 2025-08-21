package storage

type Store interface {
	Save(date []byte) error
	Load() ([]byte, error)
	GetFilename() string
}

type Storage struct {
	filename string
}

func (s *Storage) GetFilename() string {
	return s.filename
}
