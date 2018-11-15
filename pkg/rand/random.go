package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GenerateInt(data []int) {
	for i := range data {
		data[i] = rand.Intn(10)
	}
}

func GenerateString(data []string) {
	for i := range data {
		data[i] = string(letterRunes[rand.Intn(len(letterRunes))])
	}
}
