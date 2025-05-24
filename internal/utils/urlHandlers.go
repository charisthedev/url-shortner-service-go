package utils

import (
	"github.com/mattheath/base62"
)

func ObfuscateID(id int64) int64 {
    const prime = 1580030173
    return (id * prime) % 1000000007 // modulo a big prime to keep it in range
}

func ReverseObfuscateID(obfuscated int64) int64 {
    const inversePrime = 59260789 // precomputed
    const mod = 1000000007

    return (obfuscated * int64(inversePrime)) % mod
}

func HashUrl(id int64) string {
    obfuscated := ObfuscateID(id)
    return base62.EncodeInt64(obfuscated)
}