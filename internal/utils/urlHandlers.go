package utils

import "github.com/mattheath/base62"

func HashUrl (identifier int64) string {
	return base62.EncodeInt64(identifier)
}