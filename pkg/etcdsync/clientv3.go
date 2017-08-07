package etcdsync

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/job-center/server/util"
)

type etcdClientV3 struct {
	Client *clientv3.Client
	Ctx    context.Context
}

func NewCli(machines []string) *etcdClientV3 {
	ctx := context.TODO()
	cfg := clientv3.Config{
		Endpoints:   machines,
		DialTimeout: time.Second}
	cl, err := clientv3.New(cfg)
	if err != nil {
		util.Logger.SetFatal("connect to etcdsync err: ", err)
	}
	return &etcdClientV3{
		Client: cl,
		Ctx:    ctx}
}

// Set sets value to key in etcdsync server.
func (ec *etcdClientV3) Set(key, value string) error {
	_, err := ec.Client.KV.Put(ec.Ctx, key, value, nil)
	return err
}

// Get gets value from key in etcdsync server.
func (ec *etcdClientV3) Get(key string) ([]*mvccpb.KeyValue, error) {
	resp, err := ec.Client.KV.Get(ec.Ctx, key, nil)
	if err != nil {
		return nil, err
	}
	return resp.Kvs, nil
}

// CreateWatcher creates a watcher to watch a dir.
func (ec *etcdClientV3) CreateWatcher(dir string) {
	rch := ec.Client.Watch(context.Background(), "", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Println(string(ev.Type), string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}
}
