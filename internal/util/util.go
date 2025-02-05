package util

import "strconv"

func ToInt64(t any) int64 {
	switch v := t.(type) {
	case string:
		return Must(strconv.ParseInt(v, 10, 64))
	}
	panic("unknown type")
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
