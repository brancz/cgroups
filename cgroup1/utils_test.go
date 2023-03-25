package cgroup1

import (
	"bytes"
	"testing"
)

func BenchmarkReaduint64(b *testing.B) {
	b.ReportAllocs()

	buf := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		_, err := readUint("testdata/uint", buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}
