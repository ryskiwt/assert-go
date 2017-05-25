// +build release

package assert

func Do(b bool) {}

func WithMsg(b bool, v ...interface{}) {}

func WithMsgf(b bool, format string, v ...interface{}) {}
