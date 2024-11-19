package main

import (
	"fmt"
	"time"
)

// TimeError — тип для хранения времени и текста ошибки.
type TimeError struct {
	Time time.Time
	Text string
}

// Error добавляет поддержку интерфейса error для типа TimeError.
func (te TimeError) Error() string {
	return fmt.Sprintf("%v: %v", te.Time.Format(`2006/01/02 15:04:05`), te.Text)
}

// NewTimeError возвращает переменную типа TimeError c текущим временем.
func NewTimeError(text string) TimeError {
	return TimeError{
		Time: time.Now(),
		Text: text,
	}
}

func testFunc(i int) error {
	// несмотря на то что NewTimeError возвращает тип TimeError,
	// у testFunc тип возвращаемого значения равен error
	if i == 0 {
		return NewTimeError((`параметр в testFunc равен 0`))
	}
	return nil
}

func main() {
	if err := testFunc(0); err != nil {
		fmt.Println(err)
	}
}
