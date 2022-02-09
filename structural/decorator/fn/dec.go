package fn

import "strings"

type OriginalFn func([]string) string

func Process(in []string) string {
	buf := "|"
	for _, v := range in {
		buf += v + "|"
	}

	return buf
}

func ProcessWithColon(fn OriginalFn) OriginalFn {
	return func(in []string) string {
		res := fn(in)

		return strings.ReplaceAll(res, "|", ":")
	}
}
