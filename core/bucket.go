package core

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Bucket struct {
	bucket     *cache.Cache
	deleteTime time.Duration
}

func NewBucket(deleteTime time.Duration) *Bucket {
	bucket := cache.New(deleteTime*time.Minute, deleteTime*time.Minute)
	return &Bucket{deleteTime: deleteTime, bucket: bucket}
}

func (b *Bucket) AddBucket(id string, value any, d time.Duration) {
	b.bucket.Set(id, value, d)
}

func (b *Bucket) DeleteBucket(id string) {
	b.bucket.Delete(id)
}

func (b *Bucket) IsExist(id string) bool {
	_, ok := b.bucket.Get(id)
	return ok
}
