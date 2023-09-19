package main

import (
	"os"
	"time"

	"github.com/wuziqiu666/learngo/math/clockface"
)

func main() {
	time := time.Now()
	clockface.SVGWriter(os.Stdout, time)
}
