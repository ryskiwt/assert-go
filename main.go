package assert

import "fmt"

const DefaultMsg = "assertion error"

type W interface{}
type G interface{}

func NewMsg(got G, want W) string {
	return fmt.Sprintf("%s: want=%v, got=%v", DefaultMsg, want, got)
}
