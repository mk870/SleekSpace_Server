package property

import (
	"math/rand"
	"time"
)

func GeneratePropertyUniqueId() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000000000
	max := 9999999999999999
	randomInt := rand.Intn(max-min) + min
	return randomInt
}

func ReturnDeletedFavoriteProperties(availableFavoritePropertyIds, allFavoritePropertyIds []int) []int {
	deletedProperties := []int{}
	idsMap := make(map[int]bool)
	for i := 0; i < len(availableFavoritePropertyIds); i++ {
		idsMap[availableFavoritePropertyIds[i]] = true
	}
	for i := 0; i < len(allFavoritePropertyIds); i++ {
		if !idsMap[allFavoritePropertyIds[i]] {
			deletedProperties = append(deletedProperties, allFavoritePropertyIds[i])
		}
	}
	return deletedProperties
}
