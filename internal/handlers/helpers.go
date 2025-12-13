package handlers

// Helper functions for template logic

// NonEmpty returns the first non-empty string, or the default
func NonEmpty(value, defaultValue string) string {
	if len(value) > 0 {
		return value
	}
	return defaultValue
}

