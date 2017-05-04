package rediscache

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type TestStruct struct {
	Foo string
}

func TestRedis(t *testing.T) {
	rcache, err := NewRedisCache("redis://localhost:6379/0", RedisOptMaxIdle(5), RedisOptIdleTimeout(240*time.Second))
	require.NoError(t, err)

	require.NoError(t, rcache.SetInt("intval", 10))
	require.NoError(t, rcache.SetString("strval", "thestring"))
	require.NoError(t, rcache.SetBool("boolval", true))
	require.NoError(t, rcache.SetFloat("floatval", 3.14))
	require.NoError(t, rcache.SetJSON("structval", &TestStruct{
		Foo: "bar",
	}))

	intval, err := rcache.GetInt("intval")
	require.NoError(t, err)
	require.Equal(t, 10, intval)

	strval, err := rcache.GetString("strval")
	require.NoError(t, err)
	require.Equal(t, "thestring", strval)

	boolval, err := rcache.GetBool("boolval")
	require.NoError(t, err)
	require.Equal(t, true, boolval)

	floatval, err := rcache.GetFloat("floatval")
	require.NoError(t, err)
	require.Equal(t, 3.14, floatval)

	dest := new(TestStruct)
	require.NoError(t, rcache.GetJSON("structval", dest))
	require.Equal(t, "bar", dest.Foo)
}

func TestRedisExpiry(t *testing.T) {
	rcache, err := NewRedisCache("redis://localhost:6379/0", RedisOptMaxIdle(5), RedisOptIdleTimeout(240*time.Second))
	require.NoError(t, err)

	require.NoError(t, rcache.SetInt("intval", 10, Expiry(1*time.Second)))
	intval, err := rcache.GetInt("intval")
	require.NoError(t, err)
	require.Equal(t, 10, intval)

	time.Sleep(1 * time.Second)

	_, err = rcache.GetInt("intval")
	require.Equal(t, ErrCacheMiss, err)
}
