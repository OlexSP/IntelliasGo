package main

import (
	"errors"
	"fmt"
	"strconv"
)

type IntVal int

func (iv IntVal) String() {

}

func main() {

	fmt.Printf(string(1))

	inputs := []any{
		"hello",
		123,
		true,
		3.14,
	}

	var resVal string
	var err error

	for _, inputValue := range inputs {
		resVal, err = AssertStringOkNotation(inputValue)
		fmt.Printf("\n Assertion function: %s\n", "AsserInterfaceVarying")
		fmt.Printf("Input type :%[1]T | inputValue: %[1]v\n", inputValue)
	}

}

func AssertStringOkNotation(v interface{}) (string, error) {
	stringer, ok := v.(fmt.Stringer)
	if !ok {
		return "", errors.New("not supported type")
	}
	return stringer.String(), nil
}

func AssertStringVarying(v interface{}) (string, error) {
	switch convertedVal := v.(type) {
	case fmt.Stringer:
		return convertedVal.String(), nil
	case int:
		return strconv.Itoa(convertedVal), nil
	case bool:
		return strconv.FormatBool(convertedVal), nil
	default:
		return "", errors.New("not supported type")
	}
}
