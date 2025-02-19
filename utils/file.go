package utils

import "os"

func FileExists(filepath string) bool {
	if stat, err := os.Stat(filepath); err != nil {
		return stat.Mode().IsRegular()
	}
	return false
}
