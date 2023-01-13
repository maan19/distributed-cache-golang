package cache

import "time"

type Cacher interface {
	Get([]byte) ([]byte, error)
	Has([]byte) bool
	Set([]byte, []byte, time.Duration) error
	Delete([]byte) error
}
