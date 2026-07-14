package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start time.Time
	ticks []time.Duration
}

func (sw *Stopwatch) Start() {
	sw.start = time.Now()
	sw.ticks = nil
}

func (sw *Stopwatch) SaveSplit() {
	sw.ticks = append(sw.ticks, time.Now().Sub(sw.start))
}

func (sw *Stopwatch) GetResults() []time.Duration {
	return sw.ticks
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
