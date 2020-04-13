package main

var testData = ""

type mockFileSystem struct {
	mockGetDataFromFileSystem func() []byte
	mockSaveDataToFileSystem  func() []byte
}

func (m *mockFileSystem) getDataFromFileSystem() []byte {
	if m.mockGetDataFromFileSystem != nil {
		return []byte(testData)
	}
	return []byte("123")
}

func (m *mockFileSystem) saveDataToFileSystem(data []byte) {
	if m.mockSaveDataToFileSystem != nil {

	}

}
