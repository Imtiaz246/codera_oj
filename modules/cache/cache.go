package cache

type Cache[VT any] interface {
	Get(key string) (*VT, error)
	Set(key string, value VT) error
	Remove(key string)
	Flush()
}
