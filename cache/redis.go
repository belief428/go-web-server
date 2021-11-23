package logic

import (
	"time"

	"errors"

	"github.com/go-redis/redis"
)

type Redis struct {
	*RedisOption
}

type RedisOption struct {
	Addr         string
	Password     string
	DB           int
	MinIdleConns int
	IdleTimeout  int
}

var RedisClient *redis.Client

// Set
func (this *Redis) Set(key string, value interface{}, expiration int) error {
	return RedisClient.Set(key, value, time.Duration(expiration)*time.Second).Err()
}

// Get
func (this *Redis) Get(key string) (string, error) {
	return RedisClient.Get(key).Result()
}

// Del
func (this *Redis) Del(key string) error {
	return RedisClient.Del(key).Err()
}

func (this *Redis) ZAdd(key string, members ...*ScoreParams) error {
	redisZ := make([]redis.Z, 0)

	for _, v := range members {
		redisZ = append(redisZ, redis.Z{Score: v.Score, Member: v.Member})
	}
	return RedisClient.ZAdd(key, redisZ...).Err()
}

func (this *Redis) ZRangebyscore(key string, opt *ScoreRangeBy) ([]string, error) {
	return RedisClient.ZRangeByScore(key, redis.ZRangeBy{Min: opt.Min, Max: opt.Max, Offset: opt.Offset, Count: opt.Count}).Result()
}

func (this *Redis) ZRem(key string, members ...interface{}) error {
	return RedisClient.ZRem(key, members...).Err()
}

// HExists HASH Exist
func (this *Redis) HExists(key, field string) (bool, error) {
	return RedisClient.HExists(key, field).Result()
}

// HSet HASH Set
func (this *Redis) HSet(key, field string, value interface{}) error {
	return RedisClient.HSet(key, field, value).Err()
}

// HGet Hash Get
func (this *Redis) HGet(key, field string) (string, error) {
	return RedisClient.HGet(key, field).Result()
}

// HDel HASH Del
func (this *Redis) HDel(key string, fields ...string) error {
	return RedisClient.HDel(key, fields...).Err()
}

// SAdd
func (this *Redis) SAdd(key string, members ...interface{}) error {
	return RedisClient.SAdd(key, members...).Err()
}

// SIsMember
func (this *Redis) SIsMember(key string, members interface{}) (bool, error) {
	return RedisClient.SIsMember(key, members).Result()
}

// SRem
func (this *Redis) SRem(key string, members ...interface{}) error {
	return RedisClient.SRem(key, members...).Err()
}

// Run 开启
func (this *Redis) Run() error {
	option := &redis.Options{
		Network:      "",
		Addr:         this.Addr,
		Password:     this.Password,
		DB:           this.DB,
		MinIdleConns: this.MinIdleConns,
		IdleTimeout:  time.Duration(this.IdleTimeout),
	}
	RedisClient = redis.NewClient(option)

	ping, err := RedisClient.Ping().Result()

	if err != nil {
		return errors.New(ping + err.Error())
	}
	return nil
}

func NewRedis(option *RedisOption) *Redis {
	return &Redis{option}
}
