package cache

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type DelFunc func(interface{})

type (
	cacheItem struct {
		value  interface{}
		expire time.Time
	}

	Cache struct {
		io.Closer

		Duration       time.Duration
		TickerDuration time.Duration

		data    map[string]*cacheItem
		delFunc DelFunc
		lock    sync.RWMutex
		once    sync.Once
	}
)

func New(duration time.Duration, tickerDuration time.Duration) *Cache {
	return &Cache{
		Duration:       duration,
		TickerDuration: tickerDuration,
	}
}

func (c *Cache) cleaner() {
	go func() {
		for {
			c.lock.RLock()
			duration := c.TickerDuration
			c.lock.RUnlock()

			if duration == 0 {
				break
			}

			time.Sleep(duration)

			now := time.Now()
			c.lock.Lock()

			for key, item := range c.data {
				if item.expire.Before(now) {
					c.deleteValue(key, item)
				}
			}

			c.lock.Unlock()
		}
	}()
}

func (c *Cache) Values() map[string]interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	now := time.Now()

	retval := make(map[string]interface{}, len(c.data))
	for k, v := range c.data {

		if v.Valid(now) {
			retval[k] = v.value
		}
	}

	return retval
}

func (c *Cache) SetDelFunc(f DelFunc) {
	c.delFunc = f
}

func (c *Cache) Print() string {
	c.lock.RLock()
	defer c.lock.RUnlock()

	buf := bytes.NewBufferString("cache:")

	for k, v := range c.data {
		buf.WriteString(fmt.Sprintf("\n'%s':'%v' (%v)", k, v.value, v.expire.String()))
	}

	return buf.String()
}

func (c *Cache) Set(key string, value interface{}) {
	c.lock.RLock()
	duration := c.Duration
	c.lock.RUnlock()

	var expire time.Time
	if duration > 0 {
		expire = time.Now().Add(duration)
	}
	c.SetWithExpire(key, value, expire)
}

func (c *Cache) SetWithExpire(key string, value interface{}, expire time.Time) {
	c.internalInit()

	newItem := &cacheItem{value, expire}

	c.lock.Lock()
	c.data[key] = newItem
	c.lock.Unlock()
}

func (c *Cache) Get(key string) (value interface{}, exist bool) {
	c.internalInit()

	c.lock.RLock()

	item, exist := c.data[key]
	valid := exist && item.Valid(time.Now())
	if valid {
		value = item.value
	}

	c.lock.RUnlock()

	if exist && !valid {
		go c.Del(key)
	}

	exist = value != nil
	return
}

func (c *Cache) Del(key string) {
	c.internalInit()

	c.lock.Lock()
	c.deleteValue(key, nil)
	c.lock.Unlock()
}

func (c *Cache) Close() error {

	c.lock.Lock()
	defer c.lock.Unlock()

	c.TickerDuration = 0
	c.data = nil

	return nil
}

func (c *Cache) internalInit() {
	c.once.Do(func() {
		c.data = make(map[string]*cacheItem)
		c.cleaner()
	})
}

func (c *Cache) deleteValue(key string, item *cacheItem) {

	if c.delFunc != nil {
		if item == nil {
			item = c.data[key] // Only for function 'Del'
		}

		if item != nil {
			c.delFunc(item.value)
		}
	}

	delete(c.data, key)
}

func (i *cacheItem) Valid(now time.Time) bool {
	return i.expire.IsZero() || !i.expire.Before(now)
}