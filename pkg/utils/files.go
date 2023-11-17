package utils

type Files interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte) error
}
