package generators

import "io/fs"

type MockWriter struct {
	Err error
}

func (m *MockWriter) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return m.Err
}