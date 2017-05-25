package assert_test

import (
	"testing"

	"pgithub.com/ryskiwt/assert-go"
)

func BenchmarkWithMsgf_True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assert.WithMsgf(true, "%s, %s, %s, %s, %s", "0000000000", "0000000000", "0000000000", "0000000000", "0000000000")
	}
}

func BenchmarkWithMsgf_False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assert.WithMsgf(false, "%s, %s, %s, %s, %s", "0000000000", "0000000000", "0000000000", "0000000000", "0000000000")
	}
}
