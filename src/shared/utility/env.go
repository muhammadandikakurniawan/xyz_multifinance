package utility

import "os"

func GetRequiredEnv(key string) (val string) {
	val, exists := os.LookupEnv(key)
	if !exists {
		panic("REQUIRE ENV WITH KEY : " + key)
	}
	return
}

func GetEnv(key, alterVal string) (val string) {
	val, exists := os.LookupEnv(key)
	if !exists {
		val = alterVal
	}
	return
}
