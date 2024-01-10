package db

type PackagesRepository interface {
	Add(data int) (int, error)
	Read() ([]int, error)
	Remove(data int) error
}
