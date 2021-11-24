package main

import (
	"os"
	"time"

	"github.com/DreamRiveress/gfwatch"
)

func main() {
	f, err := os.Open("domain.rules")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gfw := gfwatch.New()

	if err := gfw.Decode(f); err != nil {
		println("err: ", err)
	}
	println(gfw.IsForbidden("www.baidu.com"))

	time.Sleep(time.Hour)
}