package etcdsync

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/job-center/server/util"
)

const (
	defaultTTL   = 60
	defaultTry   = 3
	deleteAction = "delete"
	expireAction = "expire"
)

// A Mutex is a mutual exclusion lock which is distributed across a cluster.
type Mutex struct {
	key    string
	id     string // The identity of the caller
	client client.Client
	kapi   client.KeysAPI
	ctx    context.Context
	ttl    time.Duration
	mutex  *sync.Mutex
}

// New creates a Mutex with the given key which must be the same
// across the cluster nodes.
// machines are the ectd cluster addresses
func LockFactory(key string, ttl int, machines []string) *Mutex {
	cfg := client.Config{
		Endpoints:               machines,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		return nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil
	}

	if len(key) == 0 || len(machines) == 0 {
		return nil
	}

	if key[0] != '/' {
		key = "/" + key
	}

	if ttl < 1 {
		ttl = defaultTTL
	}

	return &Mutex{
		key:    key,
		id:     fmt.Sprintf("%v-%v-%v", hostname, os.Getpid(), time.Now().Format("20060102-15:04:05.999999999")),
		client: c,
		kapi:   client.NewKeysAPI(c),
		ctx:    context.TODO(),
		ttl:    time.Second * time.Duration(ttl),
		mutex:  new(sync.Mutex),
	}
}

// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *Mutex) Lock() (err error) {
	m.mutex.Lock()
	for try := 1; try <= defaultTry; try++ {
		if m.lock() == nil {
			return nil
		}

		if try < defaultTry {
		}
	}
	return err
}

func (m *Mutex) lock() (err error) {
	setOptions := &client.SetOptions{
		PrevExist: client.PrevNoExist,
		TTL:       m.ttl,
	}
	resp, err := m.kapi.Set(m.ctx, m.key, m.id, setOptions)
	if err == nil {
		return nil
	}
	e, ok := err.(client.Error)
	if !ok {
		return err
	}

	if e.Code != client.ErrorCodeNodeExist {
		return err
	}

	// Get the already node's value.
	resp, err = m.kapi.Get(m.ctx, m.key, nil)
	if err != nil {
		return err
	}
	watcherOptions := &client.WatcherOptions{
		AfterIndex: resp.Index,
		Recursive:  false,
	}
	watcher := m.kapi.Watcher(m.key, watcherOptions)
	for {
		resp, err = watcher.Next(m.ctx)
		if err != nil {
			return err
		}
		if resp.Action == deleteAction || resp.Action == expireAction {
			return nil
		}
	}

}

// Unlock unlocks m.
// It is a run-time error if m is not locked on entry to Unlock.
//
// A locked Mutex is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
func (m *Mutex) Unlock() (err error) {
	defer m.mutex.Unlock()
	for i := 1; i <= defaultTry; i++ {
		var resp *client.Response
		resp, err = m.kapi.Delete(m.ctx, m.key, nil)
		if err == nil {
			return nil
		}
		util.Logger.SetDebug("Delete ", m.key, ":",resp)
		e, ok := err.(client.Error)
		if ok && e.Code == client.ErrorCodeKeyNotFound {
			return nil
		}
	}
	return err
}
