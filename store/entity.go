package store

import "store/auth"

type Data struct {
	Value string `json:"value"`
}

type Entity struct {
	Key   string     `json:"key"`
	Data  Data       `json:"data"`
	Owner *auth.User `json:"user"`
}

func NewEntity(key string, value string, user *auth.User) *Entity {
	e := Entity{
		Key: key,
		Data: Data{
			Value: value,
		},
		Owner: user,
	}
	return &e
}
