package option

import (
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

func RandomString(n int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var result []byte
	r := rand.New(src)
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Uint32()%uint32(len(bytes))])
	}
	return string(result)
}
