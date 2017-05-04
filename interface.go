package rediscache

import (
	"errors"
)

// Returned by the Get functions if the key does not exist
var ErrCacheMiss = errors.New("DispatchMe/go-rediscache: cache miss")

type Cache interface {
	SetInt(key string, val int, opts ...SetOption) error
	SetString(key string, val string, opts ...SetOption) error
	SetBool(key string, val bool, opts ...SetOption) error
	SetFloat(key string, val float64, opts ...SetOption) error

	GetInt(key string) (int, error)
	GetString(key string) (string, error)
	GetBool(key string) (bool, error)
	GetFloat(key string) (float64, error)

	SetJSON(key string, val interface{}, opts ...SetOption) error
	GetJSON(key string, dest interface{}) error
}

type SetOption func(args []interface{}) []interface{}
