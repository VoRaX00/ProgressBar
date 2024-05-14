package main

import (
	"time"

	"main.go/progressbar"
)

func main() {
	var bar progressbar.ProgressBar
	bar.NewOption(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Second * 1)
		bar.Play(int64(i))
	}
	bar.Finish()
}
