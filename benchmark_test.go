package assert_test

import (
	"testing"

	"../assert-go"
)

func BenchmarkNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		assert.Equals(v, 1)
	}
}

func BenchmarkWithMsg_True(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.WithMsg(v == 1, "assertion error")
	}
}

func BenchmarkWithMsgf_True(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.WithMsgf(v == 1, "assertion error: want=%v, got=%v", 1, v)
	}
}

func BenchmarkDo_False(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.Do(v == 0)
	}
}

func BenchmarkEquals_False(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.Equals(v, 0)
	}
}

func BenchmarkWithMsg_False(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.WithMsg(v == 0, "assertion error")
	}
}

func BenchmarkWithMsgf_False(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		assert.WithMsgf(v == 0, "assertion error: want=%v, got=%v", 1, v)
	}
}
