package database

type Config interface {
	ConnectionString() string
}

type SqliteConfig struct {
	Path     string
	InMemory bool
}

func (s SqliteConfig) ConnectionString() string {
	if s.InMemory {
		return "file::memory:?_fk=1"
	}
	return "file:" + s.Path + "?_fk=1"
}
