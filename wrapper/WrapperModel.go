package wrapper

type WrapperModel struct {
	Error 	interface{} 			`json:"error"`
	Data	map[string]interface{} 	`json:"data"`
}

func (m *WrapperModel) Init() *WrapperModel {
	return &WrapperModel{Error:"", Data:make(map[string]interface{})}
}

func (m *WrapperModel) Set(key string, value interface{}, error interface{}) {
	m.Data[key] = value
	m.Error = error
}

func (m *WrapperModel) Get() (interface{}, interface{}) {
	return m.Data, m.Error
}
