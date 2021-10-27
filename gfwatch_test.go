package gfwatch

import (
	"log"
	"os"
	"testing"
)

var (
	gfw *GfWatch
)

func TestJudge(t *testing.T) {
	if gfw.IsForbidden("www.baidu.com") {
		t.Error("judge false")
	}

	if !gfw.IsForbidden("www.google.com") {
		t.Error("judge true")
	}
}

func BenchmarkJudgeFalse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gfw.IsForbidden("www.baidu.com")
	}
}

func BenchmarkJudgeTrue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gfw.IsForbidden("www.google.com")
	}
}

func init() {
	f, err := os.Open("example/domain.rules")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gfw = New()

	if err := gfw.Decode(f); err != nil {
		log.Fatal(err)
	}
}
