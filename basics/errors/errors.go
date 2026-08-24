package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivisionByZero = errors.New("Division by zero")
var ErrNumTooLarge = errors.New("Number too large")

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (e OpError) Error() string {
	return e.Message
}

func NewOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

func doSomething() error {
	return NewOpError("doSomething", 429, "too many requests", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if a > 1000 {
		return 0, ErrNumTooLarge
	}
	return a / b, nil
}

func main() {

	val, err := divide(1001, 5)
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Println("Division by zero")
		} else if errors.Is(err, ErrNumTooLarge) {
			fmt.Println("Number too large")
		}
		return
	}

	fmt.Println(val)
}
