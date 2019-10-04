package middlewares

import "math/rand"

// Middleware package to Create a Key (6 letter unique ID for shortened URL)

func CreateKey(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	byteArray := make([]byte, length)
	for i := range byteArray {
		byteArray[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(byteArray)
}
