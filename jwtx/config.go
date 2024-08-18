package jwtx

type Config struct {
	Name         string `json:"Name"`
	AccessSecret string `json:"AccessSecret"`
	AccessExpire int    `json:"AccessExpire"`
}
