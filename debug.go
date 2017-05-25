// +build !release

package assert

import "fmt"

func Do(ok bool) {
	if !ok {
		panic("assertion error")
	}
}

func WithMsg(ok bool, v ...interface{}) {
	if !ok {
		panic("assertion error: " + fmt.Sprint(v...))
	}
}

func WithMsgf(ok bool, format string, v ...interface{}) {
	if !ok {
		panic("assertion error: " + fmt.Sprintf(format, v...))
	}
}
