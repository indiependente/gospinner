package main

import (
	"os"
	"time"

	"github.com/indiependente/gospin/spinner"
)

func main() {
	go spinner.Spin(os.Stdout, 100*time.Millisecond)
	<-time.NewTimer(2 * time.Second).C
}
