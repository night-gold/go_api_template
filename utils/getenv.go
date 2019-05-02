package utils

import "os"

// GetEnv get the env and return the associated value if it exists or return the default (df) value.
func GetEnv(env, df string) string{
	v, ok := os.LookupEnv(env)
	if ok {
		return v
	}
	return df
}