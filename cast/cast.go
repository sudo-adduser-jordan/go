package cast

import (
	"fmt"
	"strconv"
)

func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int:
		str = strconv.Itoa(arg.(int))
	case byte:
		b := arg.(byte)
		str = string(rune(b))
	case rune:
		str = string(arg.(rune))
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}

func SliceToInt(arg interface{}) []int {
	var slice []int
	switch arg.(type) {
	case []byte:
		for _, b := range arg.([]byte) {
			value, err := strconv.Atoi(string(b))
			if err != nil {
				panic("error converting []byte to []int " + err.Error())
			}
			slice = append(slice, value)
		}
	case []string:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	case []rune:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return slice
}
