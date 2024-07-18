package utils

func ReplaceStrIfNotEmpty(currentVal *string, newVal string) {
	if newVal != "" {
		*currentVal = newVal
	}
}

func ReplaceNumIfNotEmpty(currentVal *float64, newVal float64) {
	if newVal != 0 {
		*currentVal = newVal
	}
}
