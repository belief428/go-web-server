package logic

type ICache interface {
	Set(key string, value interface{}, expiration int) error
	Get(key string) (string, error)
	Del(key string) error
	ZAdd(key string, members ...*ScoreParams) error
	ZRangebyscore(key string, opt *ScoreRangeBy) ([]string, error)
	ZRem(key string, members ...interface{}) error
	HExists(key, field string) (bool, error)
	HSet(key, field string, value interface{}) error
	HGet(key, field string) (string, error)
	HDel(key string, fields ...string) error
	SAdd(key string, members ...interface{}) error
	SIsMember(key string, members interface{}) (bool, error)
	SRem(key string, members ...interface{}) error
	Run() error
}

type ScoreParams struct {
	Score  float64
	Member interface{}
}

type ScoreRangeBy struct {
	Min, Max      string
	Offset, Count int64
}
