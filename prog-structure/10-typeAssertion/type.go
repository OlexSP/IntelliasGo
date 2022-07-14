package main

import (
	"errors"
	"fmt"
	"strconv"
)

type IntVal int

func (iv IntVal) String() string {
	return strconv.Itoa(int(iv))
}

func main() {

	fmt.Println(string(100)) // dost not convert to "100",  outputs "d" unicode char

	inputs := []any{
		"hello",
		123,
		true,
		3.14,
		IntVal(100),
	}

	var resVal string
	var err error

	for _, inputValue := range inputs {
		resVal, err = AssertInterfaceVarying(inputValue)
		fmt.Printf("\n\n Assertion function: %s\n", "AsserInterfaceVarying")
		fmt.Printf("Input type :%[1]T | inputValue: %[1]v\n", inputValue)
		fmt.Printf("output type :%[1]T | outputValue: %[1]v\n", resVal)
		fmt.Printf("error value: %v\n", err)

	}

}

func AssertStringOkNotation(v interface{}) (string, error) {
	stringer, ok := v.(fmt.Stringer)
	if !ok {
		return "", errors.New("not supported type")
	}
	return stringer.String(), nil
}

func AssertInterfaceVarying(v interface{}) (string, error) {
	switch convertedVal := v.(type) {
	case fmt.Stringer:
		return convertedVal.String(), nil
	case int:
		return strconv.Itoa(convertedVal), nil
	case bool:
		return strconv.FormatBool(convertedVal), nil
	default:
		return "--", errors.New("not supported type")
	}
}
