package anwser

func Pipe(value interface{}, args ...func(interface{}) interface{}) interface{} {
	for _, cb := range args {
		value = cb(value)
	}

	return value
}

func Increment(value interface{}) interface{} {
	switch v := value.(type) {
	case string:
		return v + "1"
	case int64:
		return v + 1
	case float64:
		return v + 1
	case nil:
		return nil

	}
	return nil
}
