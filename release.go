// +build release

package assert

func Do(ok bool) {}

func Equals(want interface{}, got interface{}) {}

func WithMsg(ok bool, v ...interface{}) {}

func WithMsgf(ok bool, format string, v ...interface{}) {}
