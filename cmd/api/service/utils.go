package service

import (
	"encoding/json"
	"os"
)

type Data struct {
	UserId int
	ID     int
	Title  string
	Body   string
}

type PayLoad struct {
	Data []Data
}

func raw() ([]Data, error) {
	r, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var payload PayLoad
	err = json.Unmarshal(r, &payload)
	if err != nil {
		return nil, err
	}
	return payload.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetByID(idx int) (any, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	if idx > len(data) {
		res := make([]string, 0)
		return res, nil
	}
	return data[idx], nil
}
