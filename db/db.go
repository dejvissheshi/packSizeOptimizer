package db

// Persister defines the methods that a persistence layer should implement
type Persister interface {
	Init() error
	Insert(data int) (int, error)
	Read() ([]int, error)
	Remove(data int) error

	Close() error
}
