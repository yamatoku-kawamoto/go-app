package database

func DebugConfig() Config {
	return &SqliteConfig{
		Path:     ":memory:",
		InMemory: true,
	}
}
