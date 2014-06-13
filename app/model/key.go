package model

import (
	"fmt"
	"github.com/jiorry/gos/lib/ssdb"
)

type IModel interface {
	LoadRecord(string, string, int) ([]string, error)
	LoadRecordItems(string, string, string, int) ([]string, error)
	DelItem(string, string) (bool, error)
	Del(string) (bool, error)
	GetItemValue(string, string) (string, error)
	GetValue(string) (string, error)
}

type KeyModel struct{}

func (k *KeyModel) LoadRecordItems(name, start, end string, limit int) ([]string, error) {
	return nil, nil
}

func (k *KeyModel) LoadRecord(start, end string, limit int) ([]string, error) {
	r, err := ssdb.Do("keys", start, end, limit)
	if err != nil {
		return make([]string, 0), fmt.Errorf("lost connection, please reconnect ssdb.")
	}

	return r.Array(), nil
}

func (k *KeyModel) GetItemValue(name, key string) (string, error) {
	return "", nil
}

func (k *KeyModel) GetValue(name string) (string, error) {
	r, err := ssdb.Do("get", name)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

func (k *KeyModel) DelItem(name, key string) (bool, error) {
	return false, nil
}

func (k *KeyModel) Del(name string) (bool, error) {
	r, err := ssdb.Do("del", name)
	if err != nil {
		return false, err
	}
	return r.Bool(), nil
}
