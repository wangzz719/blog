package main

import (
	"fmt"
	"strconv"
)

type MyError struct {
	value *int
}

func (e *MyError) Error() string {
	return "error " + strconv.Itoa(*e.value)
}

/*
An interface value is nil only if the inner value and type are both unset, (nil, nil).
In particular, a nil interface will always hold a nil type.
If we store a nil pointer of type *int inside an interface value, the inner type will be *int regardless of the value of the pointer: (*int, nil).
Such an interface value will therefore be non-nil even when the pointer inside is nil.
*/
func TestMyError1(i int) error {
	var myError *MyError
	if i == 0 {
		myError = &MyError{value: &i}
	}
	return myError
}

func TestMyError2(i int) error {
	if i == 0 {
		return &MyError{value: &i}
	}
	return nil
}

func main() {
	err := TestMyError1(0)
	fmt.Println("TestMyError1 0 err==nil", err == nil)

	err = TestMyError1(1)
	fmt.Println("TestMyError1 1 err==nil", err == nil)

	err = TestMyError2(0)
	fmt.Println("TestMyError2 0 err==nil", err == nil)

	err = TestMyError2(1)
	fmt.Println("TestMyError2 1 err==nil", err == nil)
}
