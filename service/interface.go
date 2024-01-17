package service

// PackService is an interface for the PackService
type PackService interface {
	Add(packSizes []int) error
	Delete(packSizes []int) error
	Read() ([]int, error)
	Rollback([]int) error
}
