package testutils

import "os"

type MockEnviron struct {
	environment map[string]string
}

type MockFiles struct {
	files map[string]string
}

func NewMockEnviron(environment map[string]string) MockEnviron {
	return MockEnviron{environment}
}

func NewMockFiles(files map[string]string) *MockFiles {
	return &MockFiles{files}
}

func (e MockEnviron) LookupEnv(env string) (string, bool) {
	value, exists := e.environment[env]
	return value, exists
}

func (f MockFiles) ReadFile(filename string) ([]byte, error) {
	contents, exists := f.files[filename]
	if !exists {
		return []byte{}, os.ErrNotExist
	}

	return []byte(contents), nil
}

func (f *MockFiles) WriteFile(filename string, data []byte) error {
	f.files[filename] = string(data)
	return nil
}
