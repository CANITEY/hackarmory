package helpers


func StoMBool[T comparable] (slice []T) map[T]bool {
	output := map[T]bool{}
	for _, element := range slice {
		output[element] = false
	}
	return output
}


func StoMStr[T comparable] (slice []T) map[T]string {
	output := map[T]string{}
	for _, element := range slice {
		output[element] = ""
	}
	return output
}

func StoMInt[T comparable] (slice []T) map[T]int {
	output := map[T]int{}
	for _, element := range slice {
		output[element] = 0
	}
	return output
}
