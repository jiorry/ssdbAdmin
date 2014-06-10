package model

import (
	"fmt"
	"github.com/jiorry/lib/ssdb"
)

type ZsetItemModel struct{}

func (h *ZsetItemModel) LoadZsetItems(name, keyStart, start, end string, limit int) ([]string, error) {
	r, err := ssdb.Do("zkeys", name, keyStart, start, end, limit)
	if err != nil {
		return make([]string, 0), fmt.Errorf("lost connection, please reconnect ssdb.")
	}

	return r.Array(), nil
}

func (h *ZsetItemModel) LoadRecordItems(name, start, end string, limit int) ([]string, error) {
	return nil, nil
}

func (h *ZsetItemModel) LoadRecord(start, end string, limit int) ([]string, error) {
	return nil, nil
}

func (h *ZsetItemModel) GetValue(name string) (string, error) {
	return "", nil
}

func (h *ZsetItemModel) GetItemValue(name, key string) (string, error) {
	r, err := ssdb.Do("zget", name, key)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

func (h *ZsetItemModel) DelItem(name, key string) (bool, error) {
	r, err := ssdb.Do("zdel", name, key)
	if err != nil {
		return false, err
	}
	return r.Bool(), nil
}

func (h *ZsetItemModel) Del(name string) (bool, error) {
	r, err := ssdb.Do("zclear", name)
	if err != nil {
		return false, err
	}
	return r.Bool(), nil
}
