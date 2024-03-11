package entities

// Fields - поля c сообщениями об ошибке.
type Fields []Field

// Field - поле c сообщением об ошибке.
type Field struct {
	Key     string `json:"key"     yaml:"Key"     xml:"Key"     description:"Ключ."`      // Ключ
	Message string `json:"message" yaml:"Message" xml:"Message" description:"Сообщение."` // Сообщение
}

// Get - получение поля с сообщением.
func (fields Fields) Get(key string) (field Field) {
	for _, f := range fields {
		if f.Key == key {
			field = f
			break
		}
	}

	return
}
