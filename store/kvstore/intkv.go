package kvstore

type IntKV struct {
	store map[string]int
}

func newIntKV() IntKV {
	return IntKV{
		store: make(map[string]int),
	}
}

func (s IntKV) Get(key string) (int, bool, error) {
	if v, ok := s.store[key]; ok {
		return v, ok, nil
	}

	return 0, false, nil
}

func (s IntKV) Set(key string, val int) error {
	s.store[key] = val
	return nil
}
