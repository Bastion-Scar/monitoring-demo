package storage

type Storage interface {
	Save(code string, url string)
	Load(code string) (string, bool)
}

type MapStorage struct {
	data map[string]string
}

func NewMapStorage() *MapStorage {
	return &MapStorage{make(map[string]string)}
}

func (m *MapStorage) Save(code string, url string) {
	m.data[code] = url
}

func (m *MapStorage) Load(code string) (string, bool) {
	url, ok := m.data[code]
	return url, ok
}
