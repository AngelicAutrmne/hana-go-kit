package log

type Field struct {
	Key   string
	Value any
}

func FieldKV(key string, value any) Field {
	return Field{Key: key, Value: value}
}
