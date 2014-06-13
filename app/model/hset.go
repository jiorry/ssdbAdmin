package model

import (
	"fmt"
	"github.com/jiorry/gos/lib/ssdb"
)

type HsetModel struct{}

func (h *HsetModel) LoadRecordItems(name, start, end string, limit int) ([]string, error) {
	return nil, nil
}

func (h *HsetModel) LoadRecord(start, end string, limit int) ([]string, error) {
	r, err := ssdb.Do("hlist", start, end, limit)
	if err != nil {
		return make([]string, 0), fmt.Errorf("lost connection, please reconnect ssdb.")
	}

	return r.Array(), nil
}

func (h *HsetModel) GetValue(name string) (string, error) {
	return "", nil
}

func (h *HsetModel) GetItemValue(name, key string) (string, error) {
	r, err := ssdb.Do("hget", name, key)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

func (h *HsetModel) DelItem(name, key string) (bool, error) {
	r, err := ssdb.Do("hdel", name, key)
	if err != nil {
		return false, err
	}
	return r.Bool(), nil
}

func (h *HsetModel) Del(name string) (bool, error) {
	r, err := ssdb.Do("hclear", name)
	if err != nil {
		return false, err
	}
	return r.Int64() > 0, nil
}
