package etcd

import (
	"fmt"
	"testing"
)

func TestEtcd_Get(t *testing.T) {
	e := NewEtcd("STUDY")
	b, err := e.Get("/root/config/common/study_go")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(b)
}
