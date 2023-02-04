package components

// Text is not null controller function
func IsEmmty(text string) bool {
	if text == "" {
		return true
	}
	if len(text) == 0 {
		return true
	}

	return false
}

// Checkbox is checked or no controller function
func IsChecked(a string) bool {
	if a == "on" {
		return false
	}
	if a == "" {
		return true
	}
	if len(a) == 0 {
		return true
	}

	return false
}
