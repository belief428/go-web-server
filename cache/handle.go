package cache

import "github.com/ipeace/go-web-server/cache/logic"

func Set(key string, value interface{}, expiration int) error {
	return client.Set(key, value, expiration)
}

func Get(key string) (string, error) {
	return client.Get(key)
}

func Del(key string) error {
	return client.Del(key)
}

func ZAdd(key string, members ...*logic.ScoreParams) error {
	return client.ZAdd(key, members...)
}

func ZRangebyscore(key string, opt *logic.ScoreRangeBy) ([]string, error) {
	return client.ZRangebyscore(key, opt)
}

func ZRem(key string, members ...interface{}) error {
	return client.ZRem(key, members...)
}

func HExists(key, field string) (bool, error) {
	return client.HExists(key, field)
}

func HSet(key, field string, value interface{}) error {
	return client.HSet(key, field, value)
}

func HGet(key, field string) (string, error) {
	return client.HGet(key, field)
}

func HDel(key string, fields ...string) error {
	return client.HDel(key, fields...)
}

func SAdd(key string, members ...interface{}) error {
	return client.SAdd(key, members...)
}

func SIsMember(key string, members interface{}) (bool, error) {
	return client.SIsMember(key, members)
}

func SRem(key string, members ...interface{}) error {
	return client.SRem(key, members...)
}
