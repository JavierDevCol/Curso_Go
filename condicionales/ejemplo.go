package condicionales

func EsWindows(os string) bool {
	if os == "windows" {
		return true
	}
	return false
}

func EsWindowsSwitch(os string) bool {
	switch os {
	case "linux":
		return false
	case "Os X.":
		return false
	default:
		return true

	}
}
