package dlog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const bufLine = 1000 //缓存1000行

type dlogWriter struct {
	w            io.WriteCloser
	buffer       chan string
	closeStartCh chan struct{}
	closeEndCh   chan struct{}
}

func NewDlogWriter(w io.WriteCloser) *dlogWriter {
	ret := new(dlogWriter)
	ret.w = w
	ret.buffer = make(chan string, bufLine)
	ret.closeStartCh = make(chan struct{})
	ret.closeEndCh = make(chan struct{})
	go ret.realWrite()
	return ret
}

func (w dlogWriter) Write(p []byte) (n int, err error) {
	count := 0
	for {
		select {
		case <-w.closeEndCh:
			_, _ = os.Stdout.WriteString(time.Now().String() + ", dlogWriter is closed\n")
			return 0, errors.New("dlogWriter closed")
		case w.buffer <- string(p):
			return len(p), nil
		case <-time.After(time.Millisecond * 20):
			count++
			str := fmt.Sprintf(time.Now().String()+",logWrite channel is full len =%v count=%d\n", len(w.buffer), count)
			_, _ = os.Stdout.WriteString(str)
		}
	}
}

func (w dlogWriter) realWrite() {
	for {
		select {
		case p := <-w.buffer:
			_, _ = w.Write([]byte(p))
		case <-w.closeStartCh:
			_ = w.Flush()
			close(w.closeEndCh)
			return
		}
	}
}

func (w dlogWriter) Flush() (err error) {
	ch := time.After(2 * time.Second)
	for {
		select {
		case <-time.After(1 * time.Second):
			return
		case <-ch:
			return
		case p := <-w.buffer:
			_, _ = w.write([]byte(p))
		}
	}
}

func (w dlogWriter) Close() error {
	_, _ = os.Stdout.WriteString(time.Now().String() + ",dlogWriter_close(w.closeStartCh)\n")
	close(w.closeStartCh)
	<-w.closeEndCh
	_, _ = os.Stdout.WriteString(time.Now().String() + ",dlogWriter_<-w.closeEndCh\n")
	err := w.w.Close()
	_, _ = os.Stdout.WriteString(time.Now().String() + ",dlogWriter_w.w.Close()\n")
	return err
}
