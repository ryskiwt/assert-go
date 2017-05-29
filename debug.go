// +build assert

package assert

const ENABLED = true

func Do(ok bool) {
	if !ok {
		panic(DefaultMsg)
	}
}

func Equals(f func() (got G, want W)) {
	got, want := f()

	if got != want {
		panic(NewMsg(got, want))
	}
}

func WithMsg(f func() (ok bool, msg string)) {
	ok, msg := f()

	if !ok {
		panic(msg)
	}
}
