/*
Напишите функцию MyCheck(input string) error, которая должна проверять строку на соответствие
следующим параметрам:
  - Строка не должна содержать цифр (found numbers).
  - Длина должна быть меньше 20 символов (line is too long).
  - Строка должна иметь два пробела (no two spaces).

В скобках приведены тексты возвращаемых ошибок.
Функция должна находить все возможные ошибки за один вызов. Результат должен содержать слайс ошибок,
по которым строка не прошла проверку, или быть nil. Подсказка: определите свой тип для слайса ошибок.
*/
package main

import (
	"errors"
	"strings"
)

type SliceError []error

func (errs SliceError) Error() string {
	var out string
	for _, err := range errs {
		out += err.Error() + ";"
	}
	return strings.TrimRight(out, `;`)
}

func MyCheck(input string) error {
	var (
		err      SliceError
		spaces   int
		hasDigit bool
	)
	if len([]rune(input)) >= 20 {
		err = SliceError{errors.New(`line is too long`)}
	}
	for _, ch := range input {
		if ch == ' ' {
			spaces++
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}
	if hasDigit {
		err = append(err, errors.New(`found numbers`))
	}
	if spaces != 2 {
		err = append(err, errors.New(`no two spaces`))
	}
	if len(err) == 0 {
		return nil
	}
	return err
}
