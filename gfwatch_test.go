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

	if !gfw.IsForbidden("www.google.com.yyytest") {
		t.Error("judge true")
	}

	if !gfw.IsForbidden("yyytest.lvv2.com") {
		t.Error("judge true")
	}

	if !gfw.IsForbidden("yyytestlvv2.com") {
		t.Error("judge true")
	}

	if !gfw.IsForbidden("yyytest.podzone.net") {
		t.Error("judge true")
	}

	if !gfw.IsForbidden("yyytestpodzone.net") {
		t.Error("judge true")
	}

	if !gfw.IsForbidden("podzone.net.yyytest") {
		t.Error("judge true")
	}

	if !gfw.IsForbidden("podzone.netyyytest") {
		t.Error("judge true")
	}

	if gfw.IsForbidden("yyytest.podzone.net.yyytest") {
		t.Error("judge false")
	}

	if gfw.IsForbidden("odzone.net") {
		t.Error("judge false")
	}

	if !gfw.IsForbidden("e12.whatsapp.net") {
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
		gfw.IsForbidden("www.google.com.yyytest")
	}
}

func BenchmarkDecode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
