package main

import (
	"fmt"
	"strconv"
)

type MyError struct {
	value *int
}

type MyStruct struct {
	Code int64
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

func TestMyError3(i int) *MyError {
	if i == 0 {
		return &MyError{value: &i}
	}
	return nil
}

func TestMyStruct() *MyStruct {
	return nil
}

func main() {
	err1 := TestMyError1(0)
	fmt.Println("TestMyError1 0 err1==nil", err1 == nil)

	err2 := TestMyError1(1)
	fmt.Println("TestMyError1 1 err2==nil", err2 == nil)

	err3 := TestMyError2(0)
	fmt.Println("TestMyError2 0 err3==nil", err3 == nil)

	err4 := TestMyError2(1)
	fmt.Println("TestMyError2 1 err4==nil", err4 == nil)

	err5 := TestMyError3(0)
	fmt.Println("TestMyError3 0 err5==nil", err5 == nil)

	err6 := TestMyError3(1)
	fmt.Println("TestMyError3 1 err6==nil", err6 == nil)

	s := TestMyStruct()
	fmt.Println("TestMyStruct s==nil", s == nil)
}
