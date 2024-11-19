package main

import (
	"errors"
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

	/*
		лучше применить функцию As пакета errors, так как она,
		в отличие от type assertion, работает с «обёрнутыми» ошибками,
		которые разберём ниже. As находит первую в цепочке ошибку err,
		устанавливает тип, равным этому значению ошибки, и возвращает true
	*/
	if err := testFunc(0); err != nil {
		var te TimeError
		if ok := errors.As(err, &te); ok { //  Сравниваем полученную и контрольную ошибки. Сравнение идёт по типу ошибки.
			fmt.Println(te.Time, te.Text)
		} else {
			fmt.Println(err)
		}
	}
}
