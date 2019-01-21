package models

type MockSession struct{}

func (s *MockSession) DB(name string) *MockDatabase {
	return &MockDatabase{}
}

type MockDatabase struct{}

func (db *MockDatabase) C(name string) *MockCollection {
	return &MockCollection{}
}

type MockCollection struct{}

func (c *MockCollection) Insert(docs ...interface{}) error {
	return nil
}
