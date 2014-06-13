package model

import (
	"fmt"
	"github.com/jiorry/gos/lib/ssdb"
)

type HsetItemModel struct{}

func (h *HsetItemModel) LoadRecordItems(name, start, end string, limit int) ([]string, error) {
	r, err := ssdb.Do("hkeys", name, start, end, limit)
	if err != nil {
		return make([]string, 0), fmt.Errorf("lost connection, please reconnect ssdb.")
	}

	return r.Array(), nil
}

func (h *HsetItemModel) LoadRecord(start, end string, limit int) ([]string, error) {
	return nil, nil
}

func (h *HsetItemModel) GetValue(name string) (string, error) {
	return "", nil
}

func (h *HsetItemModel) GetItemValue(name, key string) (string, error) {
	r, err := ssdb.Do("hget", name, key)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

func (h *HsetItemModel) DelItem(name, key string) (bool, error) {
	r, err := ssdb.Do("hdel", name, key)
	if err != nil {
		return false, err
	}
	return r.Bool(), nil
}

func (h *HsetItemModel) Del(name string) (bool, error) {
	r, err := ssdb.Do("hclear", name)
	if err != nil {
		return false, err
	}
	return r.Bool(), nil
}
