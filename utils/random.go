package utils

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

func GenerateRandInt(data []int) {
	for i := range data {
		data[i] = rand.Intn(10)
	}
}

func GenerateRandString(data []string) {
	for i := range data {
		data[i] = string(letterRunes[rand.Intn(len(letterRunes))])
	}
}
