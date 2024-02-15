package kvstore

type StringKV struct {
	store map[string]string
}

func newStringKV() StringKV {
	return StringKV{
		store: make(map[string]string),
	}
}

func (s StringKV) Get(key string) (string, bool, error) {
	if v, ok := s.store[key]; ok {
		return v, ok, nil
	}

	return "", false, nil
}

func (s StringKV) Set(key string, val string) error {
	s.store[key] = val
	return nil
}
