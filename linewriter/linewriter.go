package linewriter

import (
	"sync"

	"github.com/popmedic/go-slices/slices"
)

type LineWriter struct {
	*slices.MaxSlice
	buffer []byte
	lock   sync.RWMutex
}

func NewLineWriter(max int) *LineWriter {
	return &LineWriter{
		slices.NewMaxSlice(max),
		[]byte{},
		sync.RWMutex{},
	}
}

func (w *LineWriter) Write(p []byte) (n int, err error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	o, n, i := len(w.buffer), len(p), 0
	buf := make([]byte, o+n)
	copy(buf[:o], w.buffer)
	for i < n {
		buf[o] = p[i]
		o++
		i++
		if string(p[i-1]) == "\n" {
			w.MaxSlice.Add(slices.NewStringItem(string(buf[:o])))
			buf = make([]byte, n-i)
			o = 0
		}
	}
	w.buffer = buf
	return n, nil
}

func (w *LineWriter) Max() int {
	return w.MaxSlice.GetMax()
}
