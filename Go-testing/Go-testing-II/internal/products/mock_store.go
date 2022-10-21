package products

type MockStorage struct {
	DataMock   []Product
	errOnWrite error
	errOnRead  error
}

func (m *MockStorage) Read(data interface{}) (err error) {
	if m.errOnRead != nil {
		return m.errOnRead
	}

	castedData := data.(*[]Product)
	*castedData = m.DataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}

	castedData := data.([]Product)
	m.DataMock = append(m.DataMock, castedData[len(castedData)-1])
	return nil
}
