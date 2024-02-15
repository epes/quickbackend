package kvstore

type KVStore[T any] interface {
	Get(key string) (T, bool, error)
	Set(key string, val T) error
}

func NewStringKV() (KVStore[string], error) {
	return newStringKV(), nil
}

func NewIntKV() (KVStore[int], error) {
	return newIntKV(), nil
}
