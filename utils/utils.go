package utils

import "math/rand"

func GetRandomInt(max int) int {
	return rand.Intn(max) + 1
}
