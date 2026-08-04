package env

import (
	"os"
	"strconv"
)

func GetEnvString(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue

}
func GetEnvInt(key string, defaultValue int) int {
	if value, exist := os.LookupEnv(key); exist {
		if intValue, err := strconv.Atoi(value); err == nil {

			return intValue
		}
	}
	return defaultValue

}
