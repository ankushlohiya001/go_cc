package go_cc

import (
	"bufio"
	"os"
	"strings"
)

type RW struct {
	rw *bufio.ReadWriter
}

func NewRW() RW {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	rw := bufio.NewReadWriter(r, w)
	return RW{rw}
}

func (self RW) Read() string {
	txt, _ := self.rw.ReadString('\n')
	return strings.TrimRight(txt, "\n")
}

func (self RW) ReadInt() int {
	txt := self.Read()
	return Int(txt)
}

func (self RW) ReadInts() (slice []int) {
	txt := self.ReadSlice()
	slice = make([]int, 0, len(txt))
	for _, v := range txt {
		slice = append(slice, Int(v))
	}
	return
}

func (self RW) ReadFloat() float64 {
	txt := self.Read()
	return Float64(txt)
}

func (self RW) ReadFloats() (slice []float64) {
	txt := self.ReadSlice()
	slice = make([]float64, 0, len(txt))
	for _, v := range txt {
		slice = append(slice, Float64(v))
	}
	return
}

func (self RW) ReadSlice() []string {
	txt := self.Read()
	return strings.Fields(txt)
}

func (self RW) Write(bts []byte) (int, error) {
	return self.rw.Write(bts)
}

func (self RW) Flush() {
	self.rw.Flush()
}
