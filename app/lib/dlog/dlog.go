package dlog

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"path"
	"runtime"
)

const defaultDir = "/home/golang/log/"
const defaultTopic = "default_topic"
const defaultHeader = `${prefix} ${level} ${time_rfc3339}`

type dLog struct {
	*log.Logger
	dw *dlogWriter
}

func NewDlog(w io.WriteCloser, topic string) *dLog {
	if len(topic) < 0 {
		topic = defaultTopic
	}
	ret := &dLog{
		Logger: log.New(topic),
	}
	ret.dw = NewDlogWriter(w)
	ret.SetOutput(ret.dw)
	ret.SetHeader(defaultHeader)
	ret.SetLevel(log.INFO)
	ret.EnableColor() //todo
	return ret
}

func (d *dLog) logStr(kv ...interface{}) string {
	_, file, line, _ := runtime.Caller(3)
	file = d.getFilePath(file)
	if len(kv)%2 != 0 {
		kv = append(kv, "unknown")
	}
	strFmt := "%s %d "
	args := []interface{}{file, line}
	for i := 0; i < len(kv); i += 2 {
		strFmt += "[%v=%+v]"
		args = append(args, kv[i], kv[i+1])
	}
	str := fmt.Sprintf(strFmt, args...)
	return str
}

func (d *dLog) getFilePath(file string) string {
	dir, base := path.Dir(file), path.Base(file)
	return path.Join(path.Base(dir), base)
}

func (d *dLog) Debug(kv ...interface{}) {
	d.Debugf("", d.logStr(kv...))
}

func (d *dLog) Info(kv ...interface{}) {
	d.Infof("", d.logStr(kv...))
}

func (d *dLog) Warn(kv ...interface{}) {
	d.Warnf("", d.logStr(kv...))
}

func (d *dLog) Error(kv ...interface{}) {
	d.Errorf("", d.logStr(kv...))
}

func (p *dLog) Close() error {
	if p.dw != nil {
		_ = p.dw.Close()
		p.dw = nil
	}
	return nil
}
