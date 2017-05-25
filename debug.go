// +build !release

package assert

import "fmt"

func Do(ok bool) {
	if !ok {
		panic("assertion error")
	}
}

func Equals(want interface{}, got interface{}) {
	if want != got {
		panic(fmt.Sprintf("assertion error: want=%v, got=%v", want, got))
	}
}

func WithMsg(ok bool, v ...interface{}) {
	if !ok {
		panic(fmt.Sprint(v...))
	}
}

func WithMsgf(ok bool, format string, v ...interface{}) {
	if !ok {
		panic(fmt.Sprintf(format, v...))
	}
}
