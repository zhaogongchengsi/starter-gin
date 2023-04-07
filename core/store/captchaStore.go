package store

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type CaptchaBucket struct {
	cache      *cache.Cache
	expiration time.Duration
}

type CaptchaItem struct {
	id, value   string
	IssueTime   time.Time
	FailureTime time.Time
}

func NewCaptchaItem(id string, value string, failureTime time.Duration) *CaptchaItem {
	now := time.Now()
	return &CaptchaItem{id: id, value: value, IssueTime: now, FailureTime: now.Add(failureTime * time.Minute)}
}

func NewCaptchaBucket() *CaptchaBucket {
	t := time.Duration(5)
	return &CaptchaBucket{cache: cache.New(t*time.Minute, t*time.Minute), expiration: t}
}

func (c *CaptchaBucket) Set(id string, value string) error {
	item := NewCaptchaItem(id, value, c.expiration)
	c.cache.Set(id, item, c.expiration*time.Minute)
	return nil
}

func (c *CaptchaBucket) Get(id string, clear bool) string {
	value, ok := c.cache.Get(id)

	if !ok {
		return ""
	}
	if clear {
		c.cache.Delete(id)
	}

	v, ok := value.(*CaptchaItem)

	if !ok {
		return ""
	}

	t := time.Now()

	if v.FailureTime.Unix() < t.Unix() {
		return ""
	}

	return v.value
}

func (c *CaptchaBucket) Verify(id, answer string, clear bool) bool {
	match := c.Get(id, clear) == answer
	return match
}

func (c *CaptchaBucket) Delete(id string) {
	c.cache.Delete(id)
}
