package internal

import "maps"

func ImmutableCopy(headers map[string]string) map[string]string {
	h := make(map[string]string, len(headers))
	maps.Copy(h, headers)
	return h
}
