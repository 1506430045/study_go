package etcd

import (
	"context"
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	etcd "github.com/coreos/etcd/clientv3"
)

var (
	ErrEtcdAddrNotFound = errors.New("etcd: etcd address not found in env variables")
	ErrKeyNotExists     = errors.New("etcd: key not exists")
)

const (
	defaultDialTimeout      = 1 * time.Second
	defaultTimeout          = 1 * time.Second
	defaultAutoSyncInterval = 30 * time.Second
)

type Etcd struct {
	once     sync.Once
	client   *etcd.Client
	err      error
	business string
}

func NewEtcd(business string) *Etcd {
	return &Etcd{business: business}
}

func (e *Etcd) Client() (*etcd.Client, error) {
	return e.tryInit()
}

func (e *Etcd) tryInit() (*etcd.Client, error) {
	e.once.Do(func() {
		endpoints, err := e.endpoints()
		if err != nil {
			e.err = err
			return
		}

		config := etcd.Config{
			Endpoints:        endpoints,
			Username:         "",
			Password:         "",
			DialTimeout:      defaultDialTimeout,
			AutoSyncInterval: defaultAutoSyncInterval,
		}
		e.client, err = etcd.New(config)
		if err != nil {
			e.err = err
			return
		}
	})
	return e.client, e.err
}

func (e *Etcd) Watch(key string) {

}

func (e *Etcd) Get(key string) ([]byte, error) {
	client, err := e.tryInit()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	response, err := client.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if len(response.Kvs) == 0 {
		return nil, ErrKeyNotExists
	}
	return response.Kvs[0].Value, nil
}

func (e *Etcd) endpoints() ([]string, error) {
	addr, ok := os.LookupEnv(e.business + "_ETCD_ADDR")
	if !ok || addr == "" {
		return nil, ErrEtcdAddrNotFound
	}
	addr = strings.TrimSpace(addr)
	addr = strings.TrimRight(addr, ";")
	endpoints := strings.Split(addr, ";")

	return endpoints, nil
}

func (e *Etcd) authInfo() (username, password string) {
	info := os.Getenv(e.business + "_ETCD_AUTHINFO")
	if info != "" {
		authInfo := strings.SplitN(info, "@", 2)
		if len(authInfo) == 2 {
			username = authInfo[0]
			password = authInfo[1]
		}
	}
	return
}
