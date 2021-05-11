package examples

import (
	"errors"
	"fmt"
)

func divide(arg1 int, arg2 int) (int, error) {
	if arg2 == 0 {
		return 0, errors.New("cannot divide a number by zero")
	}

	return arg1 / arg2, nil
}

type intArgError struct {
	arg     int
	message string
}

func (intArg *intArgError) Error() string {
	return fmt.Sprintf("%d - %s", intArg.arg, intArg.message)
}

func divide2(arg1 int, arg2 int) (int, error) {
	if arg2 == 0 {
		return 0, &intArgError{arg: arg2, message: "cannot divide a number by zero"}
	} else {
		return arg1 / arg2, nil
	}
}

func Errors() {
	for _, i := range [2][2]int{{6, 3}, {2, 0}} {
		if r, e := divide(i[0], i[1]); e == nil {
			fmt.Println("Result is", r)
		} else {
			fmt.Println("An error occurred", e)
		}
	}

	for _, i := range [2][2]int{{6, 3}, {2, 0}} {
		if r, e := divide2(i[0], i[1]); e == nil {
			fmt.Println("Result is", r)
		} else {
			fmt.Println("An error occurred", e)
		}
	}

	_, e := divide2(1, 0)
	if ae, ok := e.(*intArgError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	}
}
