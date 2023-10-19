package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max - min + 1) + min)
}

func GetWaterStatus(data int) string {
	if data <= 5 {
		return "aman"
	}

	if data >= 6 && data <= 8 {
		return "siaga"
	}

	return "bahaya"
}

func GetWindStatus(data int) string {
	if data <= 6 {
		return "aman"
	}

	if data >= 7 && data <= 15 {
		return "siaga"
	}

	return "bahaya"
}
