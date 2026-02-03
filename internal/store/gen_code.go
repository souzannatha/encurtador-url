package store

import "math/rand/v2"

const characters = "abcdefghijklmnopqrstuvwxyz0123456789"

func genCode() string {
	const n = 8
	byts := make([]byte, 8)
	for i := range n {
		byts[i] = characters[rand.IntN(len(characters))]
	}
	return string(byts)
}
