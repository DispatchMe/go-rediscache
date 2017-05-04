# go-rediscache
Simple interface-based Redis cache with type-helper functions and JSON support.

[View API docs](https://godoc.org/github.com/DispatchMe/go-rediscache)

Basically a thin wrapper for [garyburd/redigo](https://github.com/garyburd/redigo) that abstracts the pool instantiation and expiration, supports JSON, and is interface-based so it can be mocked in dependent packages.

## Example
```
import (
  "github.com/DispatchMe/go-rediscache"
  "log"
  "time"
)

func main() {
  redisCache, err := rediscache.NewRedisCache("redis://localhost:6379/0", rediscache.RedisOptIdleTimeout(240 * time.Second))
  if err != nil {
    log.Fatal(err)
  }

  err = redisCache.SetInt("myint", 10, rediscache.Expiry(5 * time.Second))
  if err != nil {
    log.Fatal(err)
  }

  val, err := redisCache.GetInt("myint")
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("Int: %d\n", val) // Int: 10

  time.Sleep(10 * time.Second)

  _, err = redisCache.GetInt("myint")

  log.Println(err) // "DispatchMe/go-rediscache: cache miss"
}
```
