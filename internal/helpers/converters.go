package helpers


func StoMBool[T comparable] (slice []T, defaultValue bool) map[T]bool {
	output := map[T]bool{}
	for _, element := range slice {
		output[element] = defaultValue
	}
	return output
}


func StoMStr[T comparable] (slice []T, defualtValue string) map[T]string {
	output := map[T]string{}
	for _, element := range slice {
		output[element] = defualtValue
	}
	return output
}

func StoMInt[T comparable] (slice []T, defaultValue int) map[T]int {
	output := map[T]int{}
	for _, element := range slice {
		output[element] = defaultValue
	}
	return output
}

func MtoS[T, U comparable] (m map[T]U) ([]T, []U) {
	slice1 := []T{}
	slice2 := []U{}
	for key, value := range m {
		slice1 = append(slice1, key)
		slice2 = append(slice2, value)
	}
	return slice1, slice2
}
