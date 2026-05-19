package v9

import (
	"context"
	"sync"
	"time"
	"yixiang.co/go-mall/pkg/global"

	"github.com/redis/go-redis/v9"
)

// RedisClient Redis service
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once ensure Redis client is a singleton
var once sync.Once

// Redis global Redis client using DB 1
var Redis *RedisClient

// ConnectRedis connect to Redis and set global client
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

// NewClient create a new Redis client
func NewClient(address string, password string, username string, db int) *RedisClient {

	// initialize custom RedisClient
	rds := &RedisClient{}
	// use default context
	rds.Context = context.Background()

	// initialize with redis.NewClient
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// ping connection
	err := rds.Ping()
	global.YSHOP_LOG.Error(err)

	return rds
}

// Ping ping Redis to verify connection
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set set key with expiration
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		global.YSHOP_LOG.Error(err.Error())
		return false
	}
	return true
}

// Get get value by key
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			global.YSHOP_LOG.Error(err.Error())
		}
		return ""
	}
	return result
}

// Has key exists; errors and redis.Nil return false
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			global.YSHOP_LOG.Error(err.Error())
		}
		return false
	}
	return true
}

// Del delete one or more keys
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		global.YSHOP_LOG.Error(err.Error())
		return false
	}
	return true
}

// FlushDB flush current Redis DB
func (rds RedisClient) FlushDB(keys ...string) bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		global.YSHOP_LOG.Error(err.Error())
		return false
	}
	return true
}

// Increment one arg: increment key by 1.
// two args: key and int64 increment amount.
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			global.YSHOP_LOG.Error(err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[0].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			global.YSHOP_LOG.Error(err.Error())
			return false
		}
	default:
		global.YSHOP_LOG.Error("too many parameters")
		return false
	}
	return true
}

// Decrement one arg: decrement key by 1.
// two args: key and int64 decrement amount.
func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			global.YSHOP_LOG.Error(err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[0].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			global.YSHOP_LOG.Error(err.Error())
			return false
		}
	default:
		global.YSHOP_LOG.Error("too many parameters")
		return false
	}
	return true
}
