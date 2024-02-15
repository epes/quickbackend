package api

import "context"

const CounterKey string = "KV_COUNTER"

func (a *api) Counter(ctx context.Context) (int, error) {
	val, ok, err := a.intKV.Get(CounterKey)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, nil
	}

	return val, nil
}

func (a *api) IncreaseCounter(ctx context.Context, addition int) (int, error) {
	val, ok, err := a.intKV.Get(CounterKey)
	if err != nil {
		return 0, err
	}

	if !ok {
		return addition, a.intKV.Set(CounterKey, addition)
	}

	return val + addition, a.intKV.Set(CounterKey, val+addition)
}
