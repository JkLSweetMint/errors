package entities

// Fields - поля c сообщениями об ошибке.
type Fields []Field

// Field - поле c сообщением об ошибке.
type Field struct {
	Key     string
	Message string
}
