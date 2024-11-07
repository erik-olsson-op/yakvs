package engine

import (
	"fmt"
	"github.com/erik-olsson-op/yakvs/model"
	"sync"
)

const (
	OK = "OK"
)

type ConcurrentMap struct {
	m    map[string][]byte
	rwMu sync.RWMutex
}

var cm ConcurrentMap

func init() {
	cm = ConcurrentMap{
		m: make(map[string][]byte),
	}
}

func Parse(q *model.Query) (*model.Response, error) {
	switch q.Keyword {
	case model.Keyword_SET:
		return set(q.Entry)
	case model.Keyword_GET:
		return get(q.Entry)
	case model.Keyword_DELETE:
		return del(q.Entry)
	default:
		return nil, fmt.Errorf("%v not valid keyword", q.Keyword)
	}
}

func get(e *model.Entry) (*model.Response, error) {
	cm.rwMu.RLock()
	v := cm.m[e.Key]
	cm.rwMu.RUnlock()
	return &model.Response{Value: v}, nil
}

func set(e *model.Entry) (*model.Response, error) {
	cm.rwMu.Lock()
	cm.m[e.Key] = e.Value
	cm.rwMu.Unlock()
	return &model.Response{Value: []byte(OK)}, nil
}

func del(e *model.Entry) (*model.Response, error) {
	cm.rwMu.Lock()
	delete(cm.m, e.Key)
	cm.rwMu.Unlock()
	return &model.Response{Value: []byte(OK)}, nil
}
