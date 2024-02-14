package entities

import (
	"strconv"
)

// ID - идентификатор ошибки.
type ID int

// String - получение строкового представления идентификатора ошибки.
func (id ID) String() string {
	return strconv.Itoa(int(id))
}
