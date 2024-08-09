package helpers

import (
	"fmt"
)

func mToSBool[T comparable] (slice []T) map[T]bool {
	output := map[T]bool{}
	for _, element := range slice {
		output[element] = false
	}
	return output
}

func MtoS[T, U comparable](slice []T, outType string) map[T]U {
	switch outType {
	case "bool":
		output := map[T]bool{}
		output = mToSBool(slice)
		return output
	}
}
