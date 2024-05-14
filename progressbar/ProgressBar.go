package progressbar

import "fmt"

type ProgressBar struct {
	percent int64
	current int64
	total   int64
	rate    string
	view    string
}

func (bar *ProgressBar) NewOption(start, total int64) {
	bar.current = start
	bar.total = total

	if bar.view == "" {
		bar.view = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.view
	}
}

func (bar *ProgressBar) NewOptionWithView(start, total int64, view string) {
	bar.view = view
	bar.NewOption(start, total)
}

func (bar ProgressBar) getPercent() int64 {
	return int64(float32(bar.current) / float32(bar.total) * 100)
}

func (bar *ProgressBar) Play(cur int64) {
	bar.current = cur
	last := bar.percent
	bar.percent = bar.getPercent()

	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.view
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.current, bar.total)
}

func (bar *ProgressBar) Finish() {
	fmt.Println()
}
