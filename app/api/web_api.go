package api

import (
	"../model"
	"github.com/jiorry/gos"
	"github.com/jiorry/lib/util"
)

var (
	pageSize = 100
)

type WebApi struct {
	gos.WebApi
}

func (a *WebApi) LoadRecord(args util.MapData) ([]string, error) {
	control := getModel(args.GetString("ctype"))

	r, err := control.LoadRecord(args.GetString("start"), args.GetString("end"), args.GetInt("limit"))
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (a *WebApi) GetValue(args util.MapData) (string, error) {
	control := getModel(args.GetString("ctype"))
	r, err := control.GetValue(args.GetString("name"))
	if err != nil {
		return "", err
	}
	return r, nil
}

func (a *WebApi) Del(args util.MapData) (bool, error) {
	control := getModel(args.GetString("ctype"))

	r, err := control.Del(args.GetString("name"))
	if err != nil {
		return false, err
	}
	return r, nil
}

func getModel(ctype string) model.IModel {
	switch ctype {
	case "key":
		return &model.KeyModel{}
	case "hset":
		return &model.HsetModel{}
	case "zset":
		return &model.ZsetModel{}
	}

	return nil
}
