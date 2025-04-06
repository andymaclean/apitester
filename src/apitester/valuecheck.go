package apitester

import "fmt"

type StringValueCheck interface {
	CheckVal(val *string) (bool, error)
}

type IntValueCheck interface {
	CheckVal(val int) (bool, error)
}

type StringEq struct {
	val string
}

func (chk *StringEq) CheckVal(val *string) (bool, error) {
	if *val == chk.val {
		return true, nil
	}
	return false, fmt.Errorf("string is %v not %v", *val, chk.val)
}

type IntEq struct {
	val int
}

func (chk *IntEq) CheckVal(val int) (bool, error) {
	if val == chk.val {
		return true, nil
	}
	return false, fmt.Errorf("value is %v not %v", val, chk.val)
}
