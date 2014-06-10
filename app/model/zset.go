package model

import (
	"fmt"
	"github.com/jiorry/lib/ssdb"
)

type ZsetModel struct{}

func (h *ZsetModel) LoadRecord(start, end string, limit int) ([]string, error) {
	r, err := ssdb.Do("zlist", start, end, limit)
	if err != nil {
		return make([]string, 0), fmt.Errorf("lost connection, please reconnect ssdb.")
	}

	return r.Array(), nil
}

func (h *ZsetModel) GetValue(name string) (string, error) {
	return "", nil
}

func (h *ZsetModel) GetItemValue(name, key string) (string, error) {
	r, err := ssdb.Do("zget", name, key)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

func (h *ZsetModel) DelItem(name, key string) (bool, error) {
	r, err := ssdb.Do("zdel", name, key)
	if err != nil {
		return false, err
	}
	return r.Bool(), nil
}

func (h *ZsetModel) Del(name string) (bool, error) {
	r, err := ssdb.Do("zclear", name)
	if err != nil {
		return false, err
	}
	return r.Bool(), nil
}
