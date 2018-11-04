package main

import (
	"gopkg.in/cheggaaa/pb.v1"
	"time"
)

func main() {
	bar := pb.StartNew(10001)
	bar.ShowBar = true
	bar.SetWidth(80)

	bar.Start()

	for i := 0; i <= 10000; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("End")
}
