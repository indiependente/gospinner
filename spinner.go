package spinner

import (
	"fmt"
	"io"
	"time"
)

// Spin writes an animated spinner on the writer w.
func Spin(w io.Writer, d time.Duration) {
	defer w.Write([]byte("\r \r"))
	for {
		for _, r := range []rune{'-', '\\', '|', '/'} {
			w.Write([]byte(fmt.Sprintf("\r%c", r)))
			time.Sleep(d)
		}
	}
}
