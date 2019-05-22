package wrapper

type WrapperModel struct {
	Error 	interface{} 	`json:"error"`
	Data	interface{} 	`json:"data"`
}

func (m *WrapperModel) Set(data interface{}, error interface{}) {
	m.Data = data
	m.Error = error
}

func (m *WrapperModel) Get() (interface{}, interface{}) {
	return m.Data, m.Error
}
