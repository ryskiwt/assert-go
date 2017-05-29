// +build !assert

package assert

const ENABLED = false

func Do(ok bool) {}

func Equals(func() (got G, want W)) {}

func WithMsg(func() (ok bool, msg string)) {}
