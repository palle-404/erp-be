package commons

import (
	"math"
	"math/rand"
	"strconv"
)

func GenerateRandomNumberString(length int) string {
	low := math.Pow10(length)
	high := math.Pow10(length+1) - 1
	randNum := int(low) + rand.Intn(int(high-low))
	return strconv.Itoa(randNum)
}
