package rediscache

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisCache struct {
	pool *redis.Pool
}

func handleNil(reply interface{}, err error) (interface{}, error) {
	if reply == nil && err == nil {
		return nil, ErrCacheMiss
	}
	return reply, err
}

func Expiry(expiry time.Duration) SetOption {
	return func(args []interface{}) []interface{} {
		return append(args, "EX", expiry.Seconds())
	}
}

// RedisOpt is an specific to the NewRedisCache function so we can get a cleaner declaration. It
// modifies a redis.Pool by reference.
type RedisOpt func(pool *redis.Pool)

// Set the maximum idle connections in the pool
func RedisOptMaxIdle(maxIdle int) RedisOpt {
	return func(p *redis.Pool) {
		p.MaxIdle = maxIdle
	}
}

// Set the maximum active connections in the pool
func RedisOptMaxActive(maxActive int) RedisOpt {
	return func(p *redis.Pool) {
		p.MaxActive = maxActive
	}
}

// Set the idle connection timeout for the pool
func RedisOptIdleTimeout(idleTimeout time.Duration) RedisOpt {
	return func(p *redis.Pool) {
		p.IdleTimeout = idleTimeout
	}
}

// NewRedisCache returns a new redis-based cache for the provided connection string, with zero or
// or more options to modify the redis.Pool.
//
// See https://www.iana.org/assignments/uri-schemes/prov/redis for connection string schema
func NewRedisCache(connectionString string, opts ...RedisOpt) (*RedisCache, error) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(connectionString, redis.DialConnectTimeout(1*time.Second), redis.DialReadTimeout(100*time.Millisecond))
			if err != nil {
				return nil, err
			}
			return c, nil
		},

		// If it's been more than a minute since we last used this connection, check
		// that it's still active by running a quick PING. Otherwise the connection
		// will be dropped by the pool and a new one will be added.
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	// Apply the options
	for _, opt := range opts {
		opt(pool)
	}

	// Test the connection - this will capture errors in the connection string
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		return nil, err
	}
	return &RedisCache{
		pool: pool,
	}, nil
}

func redisArgs(key string, val interface{}, opts []SetOption) []interface{} {
	args := []interface{}{key, val}

	for _, opt := range opts {
		args = opt(args)
	}

	return args
}

func (r *RedisCache) SetInt(key string, val int, opts ...SetOption) error {
	_, err := r.pool.Get().Do("SET", redisArgs(key, val, opts)...)
	return err
}

func (r *RedisCache) GetInt(key string) (int, error) {
	return redis.Int(handleNil(r.pool.Get().Do("GET", key)))
}

func (r *RedisCache) SetString(key string, val string, opts ...SetOption) error {
	_, err := r.pool.Get().Do("SET", redisArgs(key, val, opts)...)
	return err
}

func (r *RedisCache) GetString(key string) (string, error) {
	return redis.String(handleNil(r.pool.Get().Do("GET", key)))
}

func (r *RedisCache) SetBool(key string, val bool, opts ...SetOption) error {
	_, err := r.pool.Get().Do("SET", redisArgs(key, val, opts)...)
	return err
}

func (r *RedisCache) GetBool(key string) (bool, error) {
	return redis.Bool(handleNil(r.pool.Get().Do("GET", key)))
}

func (r *RedisCache) SetFloat(key string, val float64, opts ...SetOption) error {
	_, err := r.pool.Get().Do("SET", redisArgs(key, val, opts)...)
	return err
}

func (r *RedisCache) GetFloat(key string) (float64, error) {
	return redis.Float64(handleNil(r.pool.Get().Do("GET", key)))
}

func (r *RedisCache) SetJSON(key string, val interface{}, opts ...SetOption) error {
	marshaled, err := json.Marshal(val)
	if err != nil {
		return err
	}

	_, err = r.pool.Get().Do("SET", redisArgs(key, marshaled, opts)...)
	return err
}

func (r *RedisCache) GetJSON(key string, dest interface{}) error {
	raw, err := redis.Bytes(handleNil(r.pool.Get().Do("GET", key)))
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, dest)
}
