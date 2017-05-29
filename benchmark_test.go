package assert_test

import (
	"testing"

	assert "../assert-go"
)

func BenchmarkNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkFlag_True(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		if assert.ENABLED && v != 1 {
			panic(assert.NewMsg(v, 1))
		}
	}
}

func BenchmarkDo_True(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.Do(v == 1)
	}
}

func BenchmarkEquals_True(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.Equals(func() (assert.G, assert.W) { return v, 1 })
	}
}

func BenchmarkWithMsg_True(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.WithMsg(func() (bool, string) { return v == 1, assert.NewMsg(v, 1) })
	}
}

func BenchmarkDo_False(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.Do(v == 0)
	}
}

func BenchmarkFlag_False(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		if assert.ENABLED && v != 0 {
			panic(assert.NewMsg(v, 1))
		}
	}
}

func BenchmarkEquals_False(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.Equals(func() (assert.G, assert.W) { return v, 0 })
	}
}

func BenchmarkWithMsg_False(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.WithMsg(func() (bool, string) { return v == 0, assert.NewMsg(v, 0) })
	}
}
